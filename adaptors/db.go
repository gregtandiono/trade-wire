package adaptors

import (
	"github.com/jinzhu/gorm"

	// postgres driver for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConnector {()} gorm db adaptor
func DBConnector() *gorm.DB {
	db, err := gorm.Open(
		"postgres",
		"host=localhost user=postgres dbname=trade_wire sslmode=disable password=04120080090",
	)
	if err != nil {
		panic(err)
	}
	return db
}
