package dbmanager

import (
	"fmt"
	"database/sql"

	"github.com/sysdevguru/top-coin/price-srv/db"
)

// Currency struct
type Currency struct {
	Symbol string
	Price  float64
}

func (c *Currency) StorePrices() error {
	var price float64
	query := "SELECT coin_price FROM price_info WHERE coin_symbol = '" + c.Symbol + "'"
	err := db.WDB.QueryRow(query).Scan(&price)
	if err == sql.ErrNoRows {
		query = "INSERT INTO price_info (coin_price, coin_symbol) VALUES($1, $2)"
	} else {
		query = "UPDATE price_info SET coin_price = $1 WHERE coin_symbol = $2"
	}

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
