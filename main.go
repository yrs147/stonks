package main

import (
	"time"

	"github.com/yrs147/stonks/scraper"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		name, open, low, high, close := scraper.ScrapeData("https://in.investing.com/equities/google-inc-c")

		scraper.PrintData(name, open, low, high, close)

	}
}
