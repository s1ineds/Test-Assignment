package rates

import "time"

// Slice of coins
type Coins []Coin

// Object for single coin
type Coin struct {
	Id           string    `json:"id"`
	Symbol       string    `json:"symbol"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	CurrentPrice float64   `json:"current_price"`
	TotalVolume  int64     `json:"total_volume"`
	MarketCap    int64     `json:"market_cap"`
	LastUpdated  time.Time `json:"last_updated"`
}
