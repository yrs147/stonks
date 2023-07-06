package main

import (
	"time"

	"github.com/yrs147/stonks/scraper"
)

func main() {
	ticker := time.NewTicker(6 * time.Second)
	for range ticker.C {
		name, open, low, high, close := scraper.ScrapeData()

		scraper.PrintData(name, open, low, high, close)

	}
}
