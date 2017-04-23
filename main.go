package main

import (
	"fmt"

	"bitbucket.com/gregtandiono_/trade-wire/models"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// db := adaptors.DBConnector()
	dbg := adaptors.DBConnectorWithGorm()
	u := models.NewUser(
		uuid.NewV4(),
		"gregory tandiono",
		"gtandiono",
		"admin",
		[]byte("password"),
	)

	// u.Save(db)
	u.SaveWithGorm(dbg)

	fmt.Printf("great success")
}
