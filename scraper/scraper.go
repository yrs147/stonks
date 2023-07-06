package scraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
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

func RandUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func ScrapeData(url string) (string, float64, float64, float64, float64) {
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

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return name, open, low, high, close
}

func PrintData(name string, open float64, low float64, high float64, close float64) {
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
	// Open or create the CSV file
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	isEmpty := fileInfo.Size() == 0

	// Write headers if the file is empty
	if isEmpty {
		headers := []string{"Date", "Time", "Name", "Open", "Low", "High", "Close"}
		err = writer.Write(headers)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Write the data to the CSV file
	data := []string{currentDate, currentTime.Format("15:04:05"), name, strconv.FormatFloat(open, 'f', -1, 64), strconv.FormatFloat(low, 'f', -1, 64), strconv.FormatFloat(high, 'f', -1, 64), strconv.FormatFloat(close, 'f', -1, 64)}
	err = writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Flush any buffered data to the file
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

}
