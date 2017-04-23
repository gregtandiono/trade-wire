package main

import (
	"fmt"

	"bitbucket.com/gregtandiono_/trade-wire/models"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	uuid "github.com/satori/go.uuid"
)

func main() {
	db := adaptors.DBConnector()
	u := models.NewUser(
		uuid.NewV4(),
		"gregory tandiono",
		"gtandiono",
		"admin",
		"password",
	)

	u.Save(db)

	fmt.Printf("great success")
}
