package main

import (
	"log"
)

const kafkaTopic = "obu_data"

func main() {
	kafkaTransport, err := NewKafkaTransport(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}

	kafkaTransport.Start()
}
