package main

import (

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaTransport struct {
	consumer  *kafka.Consumer
	isRunning bool
}

func NewKafkaTransport(topic string) (*KafkaTransport, error) {
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
		consumer: c,
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
		if err == nil {
			logrus.Infof("received message: %s", string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// timeout is not an error, just means no message was received
			logrus.Errorf("consume error: %s", err)
			continue
		}
		// var obuData types.OBUData
		// json.Unmarshal(msg.Value, &obuData)
		// logrus.Infof("received message: %v", obuData)
	}

}
