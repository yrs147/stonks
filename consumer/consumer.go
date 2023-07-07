package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type StockData struct {
	Timestamp time.Time
	Name      string
	Open      float64
	Low       float64
	High      float64
	Close     float64
}

func main() {
	// Kafka broker address
	brokerAddress := "localhost:9092"

	// Create a reader with broker address, topic, and consumer group
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		Topic:    "google",
		GroupID:  "consumer1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		// Read a message from Kafka
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while consuming message: %v\n", err)
			continue
		}

		fmt.Printf("Received Message: %s\n", string(msg.Value))
	}
}
