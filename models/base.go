package models

import (
	"fmt"
	"trade-wire/adaptors"
)

type NewAbstractRecord struct{}

// Save is an abstract write method to db
func (r *NewAbstractRecord) Save(t string) {
	db := adaptors.DBConnector()
	defer db.Close()

	fmt.Println("WHYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")
	fmt.Println(&r)
	db.Table(t).Create(&r)
}
