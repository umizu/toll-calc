package main

import (
	"log"
)

const kafkaTopic = "obu_data"

func main() {
	calcService := NewCalculatorService()

	kafkaTransport, err := NewKafkaTransport(kafkaTopic, calcService)
	if err != nil {
		log.Fatal(err)
	}

	kafkaTransport.Start()
}
