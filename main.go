package main

import (
	"fmt"
	"log"
	"math/rand"
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

func scrapeData() {
	c := colly.NewCollector(colly.AllowedDomains("in.investing.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randUserAgent())
	})

	c.OnHTML("h1.main-title.js-main-title span.text", func(e *colly.HTMLElement) {
		index := e.Text
		index = strings.TrimSpace(index)
		fmt.Println("Name : ", index)
	})

	c.OnHTML("div.last-price-and-wildcard bdo.last-price-value.js-streamable-element", func(e *colly.HTMLElement) {
		close := e.Text
		if close != "" {
			fmt.Println("Price: ", close)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	err := c.Visit("https://in.investing.com/equities/facebook-inc")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		scrapeData()
	}
}
