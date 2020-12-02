package main

import (
	"fmt"
	"github.com/Study-of-slag-weng/go-kafka-avro"
	"github.com/bsm/sarama-cluster"
)

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topics = []string{"test"}

func main() {
	consumerCallbacks := kafka.ConsumerCallbacks{
		OnDataReceived: func(msg kafka.Message) {
			fmt.Println(msg)
		},
		OnError: func(err error) {
			fmt.Println("Consumer error", err)
		},
		OnNotification: func(notification *cluster.Notification) {
			fmt.Println(notification)
		},
	}

	consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topics, "consumer-group", consumerCallbacks, true)
	if err != nil {
		fmt.Println(err)
	}
	consumer.Consume()
}
