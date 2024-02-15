package main

import (
	"log"
)

const kafkaTopic = "obu_data"

func main() {
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaTransport, err := NewKafkaTransport(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}

	kafkaTransport.Start()
}
