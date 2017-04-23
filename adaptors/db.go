package adaptors

import (
	"database/sql"

	// pq is just used for a driver that's why it's left blank
	_ "github.com/lib/pq"
)

// DBConnector is just compartmentalized db adaptor
func DBConnector() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=postgres password=04120080090 dbname=trade_wire sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
