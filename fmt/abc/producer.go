package main

import (
	"log"
	"my-go/fmt/go-nsq"
)

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4152", config)
	if err != nil {
		log.Fatal(err)
	}

	messageBody := []byte("hello, abc")
	topicName := "tanghuamao"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer.
	producer.Stop()
}
