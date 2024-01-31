package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
	"github.com/umizu/tollcalc/types"
)

const kafkaTopic = "obudata"

func main() {
	receiver, err := NewDataReceiver()
	if err != nil {
		panic(err)
	}

	defer receiver.prod.Close()
	http.HandleFunc("/ws", receiver.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  *kafka.Producer
}

func (dr *DataReceiver) produceObuData(data types.OBUData) error {
	// Produce messages to topic (asynchronously)
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	topic := kafkaTopic
	if err = dr.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: b,
	}, nil); err != nil {
		return err
	}

	go func() {
		for e := range dr.prod.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	return nil
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("obu client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
		}
		// fmt.Printf("received OBU data from [%d]: <lat %.2f, lng %.2f>\n", data.OBUId, data.Lat, data.Lon)
		if err := dr.produceObuData(data); err != nil {
			log.Println("produce error:", err)
		}
	}
}
