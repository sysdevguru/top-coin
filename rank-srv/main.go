package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/sysdevguru/top-coin/rank-srv/dbmanager"

	"github.com/lucazulian/cryptocomparego"
)

var (
	checkInterval = 1
)

func getUpdates() {
	ctx := context.TODO()

	client := cryptocomparego.NewClient(nil)
	coinList, _, err := client.Coin.List(ctx)
	if err != nil {
		fmt.Printf("rank-srv: connecting service failure%v\n", err)
		return
	}

	for _, listing := range coinList {
		rank, _ := strconv.Atoi(listing.SortOrder)
		currency := dbmanager.Currency{
			Symbol: listing.Name,
			Rank:   rank,
		}
		currency.StoreRanks()
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
