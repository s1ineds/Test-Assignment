package main

import (
	"coin/rates"
	"coin/scraper"
	"fmt"
	"log"
	"os"
	"time"
)

// The program should run with command line arguments
func main() {
	var cc rates.CoinClient

	if len(os.Args) < 2 {
		fmt.Println("No command supplied. Exit program.")
		os.Exit(1)
	}

	if os.Args[1] == "polling" {
		go spinner()
		cc.Polling()
	}

	if os.Args[1] == "get" {
		if os.Args[2] != "" {
			cc.GetCoin(os.Args[2])
		} else {
			log.Println("Enter a coin name. Or your coin doesn't exists.")
		}
	}

	if os.Args[1] == "scrape" {
		hs := scraper.HypeScraper{}

		go spinner()

		hs.ScrapePage()
	}
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
