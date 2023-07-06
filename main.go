package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
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

type StockData struct {
	Name  string
	Close string
}

func randUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.tickertape.in"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randUserAgent())
	})

	var wg sync.WaitGroup
	dataCh := make(chan StockData)

	c.OnHTML("span.jsx-4049911629.ticker.text-teritiary.font-medium", func(e *colly.HTMLElement) {
		index := e.Text
		index = strings.TrimSpace(index)
		dataCh <- StockData{Name: index}
	})

	c.OnHTML("div.jsx-3168773259.quote-box-root.with-children span.jsx-3168773259.current-price.typography-h1.text-primary", func(e *colly.HTMLElement) {
		close := e.Text
		close = strings.ReplaceAll(close, ",", "")
		close = strings.TrimPrefix(close, "₹")
		if close != "" {
			dataCh <- StockData{Close: close}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := c.Visit("https://www.tickertape.in/indices/nifty-50-index-.NSEI")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	var stockData StockData
	for d := range dataCh {
		if d.Name != "" {
			stockData.Name = d.Name
		}
		if d.Close != "" {
					stockData.Close = d.Close
		}

		if stockData.Name != "" && stockData.Close != "" {
			fmt.Println("Name:", stockData.Name)
			fmt.Println("Close:", stockData.Close)
			// Reset stockData for the next iteration
			stockData = StockData{}
		}
	}
}
