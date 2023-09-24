package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
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
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	}
)

func RandUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func ScrapeData(url string) (string, string, float64, float64, float64, float64) {
	c := colly.NewCollector(colly.AllowedDomains("in.investing.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandUserAgent())
	})

	var (
		name  string
		open  float64
		low   float64
		high  float64
		close float64
	)

	c.OnHTML("h1.main-title.js-main-title span.text", func(e *colly.HTMLElement) {
		index := e.Text
		index = strings.TrimSpace(index)
		index = strings.ReplaceAll(index, " ", "")
		name = index
	})

	c.OnHTML("div.common-data-item dt.common-data-label span.text:contains('Open')", func(e *colly.HTMLElement) {
		openValue := e.DOM.Parent().Next().Text()
		openValue = strings.TrimSpace(openValue)
		openValue = strings.ReplaceAll(openValue, ",", "")
		if openValue != "" {
			open, _ = strconv.ParseFloat(openValue, 64)
		}
	})

	c.OnHTML("div.common-data-item dt.common-data-label span.text:contains('Day')", func(e *colly.HTMLElement) {
		daysRangeValue := e.DOM.Parent().Next().Text()
		daysRangeValue = strings.TrimSpace(daysRangeValue)
		daysRangeValue = strings.ReplaceAll(daysRangeValue, ",", "")
		splitValues := strings.Split(daysRangeValue, " - ")
		if len(splitValues) == 2 {
			low, _ = strconv.ParseFloat(splitValues[0], 64)
			high, _ = strconv.ParseFloat(splitValues[1], 64)
		}
	})

	c.OnHTML("div.last-price-and-wildcard bdo.last-price-value.js-streamable-element", func(e *colly.HTMLElement) {
		price := e.Text
		price = strings.TrimSpace(price)
		price = strings.ReplaceAll(price, ",", "")
		if price != "" {
			close, _ = strconv.ParseFloat(price, 64)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	// Load the local time zone
	loc, err := time.LoadLocation("Asia/Kolkata") 
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now().In(loc)
	timestamp := currentTime.Format("02-01-2006 15:04:05")

	return timestamp, name, open, low, high, close
}

func main() {
	ticker := time.NewTicker(4 * time.Second)
	// var stockData []StockData

	// Kafka broker address
	brokerAddress := os.Getenv("BROKER_ADDRESS2")
	url := os.Getenv("STOCK2")

	// Create a writer with broker address and topic
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   "stock2",
	})

	defer writer.Close()

	for range ticker.C {
		timestamp, name, open, low, high, close := ScrapeData(url)
		data := StockData{
			Timestamp: timestamp,
			Name:      name,
			Open:      open,
			Low:       low,
			High:      high,
			Close:     close,
		}

		// Convert data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Failed to convert to JSON: %v\n", err)
			continue
		}

		// Write the message to Kafka
		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: jsonData,
		})
		if err != nil {
			log.Printf("Failed to produce message: %v\n", err)
		}

		log.Printf("Sent Message: %s\n", jsonData)
	}

	writer.Close()
}
