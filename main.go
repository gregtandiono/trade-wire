package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
}

func dbConnector() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=postgres password=04120080090 dbname=trade_wire sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	return db
}
