package main

import (
	"fmt"
	"time"

	"github.com/sysdevguru/top-coin/price-srv/dbmanager"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

var (
	token         = "97dcc3a8-ea8b-4488-889d-6433446d3193"
	checkInterval = 1
)

func getUpdates() {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: token,
	})

	listings, err := client.Cryptocurrency.LatestListings(&cmc.ListingOptions{
		Limit: 200,
	})
	if err != nil {
		fmt.Printf("price-srv: connecting to service failure:%v\n", err)
		return
	}

	for _, listing := range listings {
		currency := dbmanager.Currency{
			Symbol: listing.Symbol,
			Price:  listing.Quote["USD"].Price,
		}
		currency.StorePrices()
	}
}

func main() {
	cha := make(chan int)
	getUpdates()

	go func() {
		checkTicker := time.NewTicker(time.Duration(checkInterval) * time.Minute)
		defer checkTicker.Stop()

		for {
			time.Sleep(time.Second)
			select {
			case <-checkTicker.C:
				getUpdates()
			}
		}
	}()
	<-cha
}
