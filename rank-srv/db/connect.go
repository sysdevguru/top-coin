package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	WDB *sql.DB
)

func init() {
	var err error
	psqlConnectInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"postgresql", 5432, "wattx", "rootroot", "coindb")
	WDB, err = sql.Open("postgres", psqlConnectInfo)
	if err != nil {
		fmt.Printf("Database connection failure:%v\n", err)
		panic(err)
	}
}
