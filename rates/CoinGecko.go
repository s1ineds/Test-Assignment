package rates

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CoinGecko struct{}

// Method that executes GET request and returns slice of coins
func (cg CoinGecko) makeGet(url string) Coins {
	defer cg.recoverFromPanic()

	var coins Coins

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	unmarshallError := json.Unmarshal(body, &coins)
	if err != nil {
		log.Fatal(unmarshallError)
	}

	return coins
}

// Method for recovering from panic if something wrong happens with get request.
func (cg *CoinGecko) recoverFromPanic() {
	if r := recover(); r != nil {
		log.Println("RECOVERED!", r)
	}
}
