package adaptors

import (
	"github.com/jinzhu/gorm"

	// postgres driver for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConnector {()} gorm db adaptor
func DBConnector() *gorm.DB {
	_, _, dbConfig := GetEnvironmentVariables()
	db, err := gorm.Open(
		"postgres",
		"host="+dbConfig["host"]+" user="+dbConfig["user"]+" dbname="+dbConfig["name"]+" sslmode=disable password="+dbConfig["password"],
	)
	if err != nil {
		panic(err)
	}
	return db
}
