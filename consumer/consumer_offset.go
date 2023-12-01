package main

// import (
// 	"context"
// 	"log"

// 	"github.com/segmentio/kafka-go"
// )

// func main() {
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{"localhost:9092"},
// 		Topic:     "user-registered",
// 		Partition: 0,
// 		MinBytes:  10e3, // 10KB
// 		MaxBytes:  10e6, // 10MB
// 	})

// 	defer r.Close()

// 	// Read and store the latest offset
// 	lastOffset := r.Offset()

// 	// Set the offset for the next read to prevent re-reading old messages
// 	if err := r.SetOffset(lastOffset + 1); err != nil {
// 		log.Printf("Error setting offset: %v", err)
// 		return
// 	}

// 	for {
// 		msg, err := r.ReadMessage(context.Background())
// 		if err != nil {
// 			log.Printf("Error reading message: %v", err)
// 			break
// 		}

// 		log.Printf("Received message: %s", string(msg.Value))
// 		// Logic to handle received messages...
// 	}
// }
