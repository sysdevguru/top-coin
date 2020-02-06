package dbmanager

import (
	"fmt"
	"database/sql"

	"github.com/sysdevguru/top-coin/rank-srv/db"
)

// Currency struct
type Currency struct {
	Symbol string
	Rank   int
}

// StoreRanks stores ranks into DB
func (c *Currency) StoreRanks() error {
	var rank int
	query := "SELECT coin_rank FROM rank_info WHERE coin_symbol = '" + c.Symbol + "'"
	err := db.WDB.QueryRow(query).Scan(&rank)
	if err == sql.ErrNoRows {
		query = "INSERT INTO rank_info (coin_rank, coin_symbol) VALUES($1, $2)"
	} else {
		query = "UPDATE rank_info SET coin_rank = $1 WHERE coin_symbol = $2"
	}

	stmt, err := db.WDB.Prepare(query)
	if err != nil {
		fmt.Printf("rank-srv: unexpected db failure:%v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Rank, c.Symbol)
	if err != nil {
		fmt.Printf("rank-srv: unexpected db failure:%v\n", err)
		return err
	}
	return nil
}
