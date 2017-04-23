package main

import (
	"bitbucket.com/gregtandiono_/trade-wire/models"

	"fmt"

	"reflect"

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
		[]byte("password"),
	)

	u.Save(db)
	allUsers := u.FetchAll(db)
	fmt.Println(reflect.TypeOf(allUsers))
}
