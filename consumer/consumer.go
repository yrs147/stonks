package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/segmentio/kafka-go"
)

type StockData struct {
	Timestamp string
	Name      string
	Open      float64
	Low       float64
	High      float64
	Close     float64
}

var (
	openGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_open",
		Help: "Open price of the stock",
	}, []string{"name"})
	lowGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_low",
		Help: "Low price of the stock",
	}, []string{"name"})
	highGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_high",
		Help: "High price of the stock",
	}, []string{"name"})
	closeGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_close",
		Help: "Close price of the stock",
	}, []string{"name"})
)

func init() {
	prometheus.MustRegister(openGauge)
	prometheus.MustRegister(lowGauge)
	prometheus.MustRegister(highGauge)
	prometheus.MustRegister(closeGauge)
}

func main() {
	// Kafka broker address
	brokerAddress1 := os.Getenv("BROKER_ADDRESS1")
	brokerAddress2 := os.Getenv("BROKER_ADDRESS2")

	// Create a reader with broker address, topic, and consumer group
	reader1 := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress1},
		Topic:    "stock1",
		GroupID:  "consumer1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// Create a reader with broker address, topic, and consumer group
	reader2 := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress2},
		Topic:    "stock2",
		GroupID:  "consumer1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// Start Prometheus HTTP server
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9010", nil))
	}()

	for {
		// Read a message from Kafka
		msg1, err := reader1.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while consuming message: %v\n", err)
			continue
		}

		processMessage(string(msg1.Value))
		defer reader1.Close()

		// Read a message from Kafka
		msg2, err := reader2.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error while consuming message: %v\n", err)
			continue
		}

		processMessage(string(msg2.Value))
		defer reader2.Close()
	}
}

func processMessage(msg string) {
	// Parse the received message
	var stockData StockData
	err := json.Unmarshal([]byte(msg), &stockData)
	if err != nil {
		log.Printf("Error while parsing JSON: %v\n", err)
		return
	}

	// Update Prometheus metrics with labels
	openGauge.WithLabelValues(stockData.Name).Set(stockData.Open)
	lowGauge.WithLabelValues(stockData.Name).Set(stockData.Low)
	highGauge.WithLabelValues(stockData.Name).Set(stockData.High)
	closeGauge.WithLabelValues(stockData.Name).Set(stockData.Close)

	fmt.Printf("Received Message: %+v\n", stockData)
}
