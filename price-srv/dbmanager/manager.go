package dbmanager

import (
	"fmt"

	"github.com/sysdevguru/top-coin/price-srv/db"
)

// Currency struct
type Currency struct {
	Symbol string
	Price  float64
}

func (c *Currency) StorePrices() error {
	query := "INSERT INTO price_info(coin_price, coin_symbol) VALUES($1, $2) ON CONFLICT ON CONSTRAINT price_symbol DO UPDATE SET coin_price = $1 WHERE price_info.coin_symbol = $2"

	stmt, err := db.WDB.Prepare(query)
	if err != nil {
		fmt.Printf("price-srv: unexpected db failure:%v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Price, c.Symbol)
	if err != nil {
		fmt.Printf("price-srv: unexpected db failure:%v\n", err)
		return err
	}
	return nil
}
