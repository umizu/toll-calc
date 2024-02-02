package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"github.com/umizu/tollcalc/types"
)

type KafkaTransport struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaTransport(topic string, svc CalculatorServicer) (*KafkaTransport, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)
	return &KafkaTransport{
		consumer:    c,
		calcService: svc,
	}, nil
}

func (c *KafkaTransport) Start() {
	logrus.Info("starting kafka transport")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaTransport) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil && !err.(kafka.Error).IsTimeout() {
			// timeout is not an error. it just means no message was received
			logrus.Errorf("consume error: %s", err)
			continue
		}

		var obuData types.OBUData
		if err := json.Unmarshal(msg.Value, &obuData); err != nil {
			logrus.Errorf("json deserialization error: %s", err)
			continue
		}

		distance, err := c.calcService.CalculateDistance(obuData)
		if err != nil {
			logrus.Errorf("error calculating distance: %s", err)
			continue
		}

		logrus.Infof("distance: %.2f", distance)
	}

}
