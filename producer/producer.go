package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Create Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user-registered",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	// Produce messages to Kafka topic
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("User registered %d: John Doe", i)
		kmsg := kafka.Message{
			Value: []byte(msg),
		}

		err := writer.WriteMessages(context.Background(), kmsg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			log.Println(msg)
		}

		time.Sleep(1 * time.Second) // Adding delay between messages
	}
}
