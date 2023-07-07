package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/yrs147/stonks/scraper"
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
	ticker := time.NewTicker(4 * time.Second)
	var stockData []StockData

	// Kafka broker address
	brokerAddress := "localhost:9092"

	// Create a writer with broker address and topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   "google",
	})

	defer writer.Close()

	for range ticker.C {
		name, open, low, high, close := scraper.ScrapeData("https://in.investing.com/equities/google-inc-c")
		data := StockData{
			Timestamp: time.Now(),
			Name:      name,
			Open:      open,
			Low:       low,
			High:      high,
			Close:     close,
		}
		stockData = append(stockData, data)

		// Convert data to a byte array or string
		value := []byte(fmt.Sprintf("%v", data))

		// Write the message to Kafka
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: value,
		})
		if err != nil {
			log.Printf("Failed to produce message: %v\n", err)
		}

		stockData = append(stockData, data)
		scraper.PrintData(name, open, low, high, close)
	}

	writer.Close()
}
