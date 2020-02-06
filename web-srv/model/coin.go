package model

import (
	"fmt"

	. "github.com/sysdevguru/top-coin/web-srv/db"
)

// Coin struct
type Coin struct {
	ID     int     `json:"-"`
	Rank   int     `json:"rank"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

// GetCoins fetches <limied> amount of top coins from DB
func (c *Coin) GetTopCoins(limit string, top string) ([]Coin, error) {
	coins := []Coin{}
	if top != "true" {
		query := "SELECT coin_rank, coin_symbol FROM rank_info ORDER BY coin_rank ASC "
		if limit != "" {
			query += "LIMIT " + limit
		} else {
			query += "LIMIT 200"
		}
	
		rows, err := WDB.Query(query)
		if err != nil {
			fmt.Printf("web-srv: db operation failure:%v\n", err)
			return coins, err
		}
		defer rows.Close()

		for rows.Next() {
			coin := Coin{}
			rows.Scan(&coin.Rank, &coin.Symbol)
			subQuery := "SELECT coin_price FROM price_info WHERE coin_symbol = '" + coin.Symbol + "'"
			WDB.QueryRow(subQuery).Scan(&coin.Price)
			coins = append(coins, coin)
		}
		return coins, nil
	}
	
	query := "SELECT r.coin_rank, r.coin_symbol, p.coin_price FROM rank_info r JOIN price_info p ON r.coin_symbol = p.coin_symbol ORDER BY r.coin_rank ASC "
	if limit != "" {
		query += "LIMIT " + limit
	} else {
		query += "LIMIT 200"
	}

	rows, err := WDB.Query(query)
	if err != nil {
		fmt.Printf("web-srv: db operation failure:%v\n", err)
		return coins, err
	}
	defer rows.Close()

	for rows.Next() {
		coin := Coin{}
		rows.Scan(&coin.Rank, &coin.Symbol, &coin.Price)
		coins = append(coins, coin)
	}

	return coins, nil
}
