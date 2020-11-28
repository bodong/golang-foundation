package main

import (
	"fmt"
	"os"
)

func main() {
	t, err := NewTrade("MSFT", 2, 4.5, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", t)
	}
}

//NewTrade will create new object of Trade
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be > 0, given %d", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be > 0, given %f", price)
	}

	trade := &Trade{
		symbol: symbol,
		volume: volume,
		price:  price,
		buy:    buy,
	}

	return trade, nil

}

//Trade object
type Trade struct {
	symbol string
	volume int
	price  float64
	buy    bool
}
