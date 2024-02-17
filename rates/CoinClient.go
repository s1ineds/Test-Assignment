package rates

import (
	"coin/config"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type CoinClient struct{}

// Method that get all coins that API returns
// and pass list of coins to the next method
func (cc *CoinClient) GetAllCoins() {
	var coinRequest RequestMaker = CoinGecko{}

	coins := coinRequest.makeGet(config.COIN_URL)

	go cc.createCSV(coins)
}

// Method that creates and fills up the csv file
func (cc *CoinClient) createCSV(records Coins) {
	if _, err := os.Stat(config.COIN_FILE); err != nil {
		os.Create(config.COIN_FILE)
	} else {
		os.Truncate(config.COIN_FILE, 0)
	}

	file, err := os.OpenFile(config.COIN_FILE, os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	var row []string = make([]string, 0)
	for _, coin := range records {
		row = append(row, coin.Id, coin.Name, coin.Symbol, coin.Image, fmt.Sprintf("%.2f", coin.CurrentPrice),
			fmt.Sprintf("%d", coin.TotalVolume), fmt.Sprintf("%d", coin.MarketCap), coin.LastUpdated.Format("2006 January 02"))

		writeErr := csvWriter.Write(row)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
		row = make([]string, 0)
	}

	csvWriter.Flush()
}

// Method that runs API polling in one time per 10 minutes
func (cc *CoinClient) Polling() {
	for {
		go cc.GetAllCoins()
		log.Println("New data saved.")
		time.Sleep(10 * time.Minute)
	}
}

// Method that output single coin by coin name or symbol
func (cc *CoinClient) GetCoin(coinname string) {
	csvFile, err := os.OpenFile(config.COIN_FILE, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	records, _ := csvReader.ReadAll()

	for _, row := range records {
		if row[0] == coinname || row[2] == coinname {
			output := fmt.Sprintf(`
			Coin name: %s
			Symbol: %s
			Image: %s
			Current Price: %s
			Total Volume: %s
			Market Cap: %s
			Last updated: %s`,
				row[1], row[2], row[3], row[4],
				row[5], row[6], row[7])

			fmt.Println(output)
		}
	}
}
