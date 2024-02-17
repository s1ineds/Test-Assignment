package scraper

import (
	"coin/config"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type HypeScraper struct{}

// Method that starts the scraping process
func (hs *HypeScraper) ScrapePage() {
	collector := colly.NewCollector()

	// When some error occurred displaying the error message
	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// Execute FetchData method when html tag was found
	collector.OnHTML("div.row__top", hs.FetchData)

	// Create a channel for sync goroutine
	var doneChan chan string = make(chan string)

	// Go to page!
	go hs.visitPage(collector, doneChan)

	// Waiting until page has visited
	log.Println(<-doneChan)
}

// Method that visits page as goroutine
func (hs *HypeScraper) visitPage(collector *colly.Collector, doneChan chan string) {
	visitErr := collector.Visit(config.SCRAPE_URL)
	if visitErr != nil {
		log.Fatal(visitErr)
	}

	// When page has been visited send message to channel
	doneChan <- "page visited"
}

// Method that trigger when the app has found a html tag
func (hs *HypeScraper) FetchData(e *colly.HTMLElement) {
	// Sigle contributor object
	co := Contributor{}

	// List of cintributor objects
	var contributorList []Contributor = make([]Contributor, 0)

	// Extract text from html tags
	co.Rank = e.ChildText("div.rank>span")
	co.Influencer = e.ChildText(".contributor__title")
	co.Category = e.ChildText("div.ellipsis")
	co.Followers = e.ChildText(".subscribers")
	co.Country = e.ChildText(".audience")
	co.Auth = e.ChildText(".authentic")
	co.Avg = e.ChildText(".engagement")

	// Add contributor object to list
	contributorList = append(contributorList, co)

	fmt.Printf("%v\n", contributorList)

	// Create the channel
	var csvChan chan string = make(chan string)

	// Execute the method as goroutine that creates and fill up csv file
	go hs.createCsv(contributorList, csvChan)

	// Wait until csv file has been filled
	log.Println(<-csvChan)
}

// Method that creates and fills csv file with data
func (hs *HypeScraper) createCsv(objects []Contributor, csvChan chan string) {
	// Open the csv file of it's exists or create a new if it doesn't.
	file, err := os.OpenFile(config.SCRAPE_DATA, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	var row []string = make([]string, 0)
	for _, obj := range objects {
		row := append(row, obj.Rank, obj.Influencer, obj.Category,
			obj.Followers, obj.Country, obj.Auth, obj.Avg)

		writeErr := csvWriter.Write(row)
		if writeErr != nil {
			log.Fatal(writeErr)
		}

		// /row = make([]string, 0)
	}

	csvWriter.Flush()

	// When the file is written send message to channel
	csvChan <- "csv flushed"
}
