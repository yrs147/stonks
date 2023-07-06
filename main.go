package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

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

func randUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func scrapeData() (string, float64, float64, float64, float64) {
	c := colly.NewCollector(colly.AllowedDomains("in.investing.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randUserAgent())
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
		name = index
	})

	c.OnHTML("div.common-data-item dt.common-data-label span.text:contains('Open')", func(e *colly.HTMLElement) {
		openValue := e.DOM.Parent().Next().Text()
		openValue = strings.TrimSpace(openValue)
		if openValue != "" {
			open, _ = strconv.ParseFloat(openValue, 64)
		}
	})

	c.OnHTML("div.common-data-item dt.common-data-label span.text:contains('Day')", func(e *colly.HTMLElement) {
		daysRangeValue := e.DOM.Parent().Next().Text()
		daysRangeValue = strings.TrimSpace(daysRangeValue)
		splitValues := strings.Split(daysRangeValue, " - ")
		if len(splitValues) == 2 {
			low, _ = strconv.ParseFloat(splitValues[0], 64)
			high, _ = strconv.ParseFloat(splitValues[1], 64)
		}
	})

	c.OnHTML("div.last-price-and-wildcard bdo.last-price-value.js-streamable-element", func(e *colly.HTMLElement) {
		price := e.Text
		price = strings.TrimSpace(price)
		if price != "" {
			close, _ = strconv.ParseFloat(price, 64)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	err := c.Visit("https://in.investing.com/equities/facebook-inc")
	if err != nil {
		log.Fatal(err)
	}

	return name, open, low, high, close
}

func printData(name string, open float64, low float64, high float64, close float64) {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	fmt.Println("Current Date:", currentDate)
	fmt.Println("Current Time:", currentTime.Format("15:04:05"))
	fmt.Println("Name:", name)
	fmt.Println("Open:", open)
	fmt.Println("Low:", low)
	fmt.Println("High:", high)
	fmt.Println("Close:", close)
	fmt.Println("------------------------")
	
}

func main() {
	ticker := time.NewTicker(6 * time.Second)
	for range ticker.C {
		name, open, low, high, close := scrapeData()

		printData(name,open,low,high,close)
		
	}
}
