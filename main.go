package main

import (
	"fmt"

	"bitbucket.com/gregtandiono_/trade-wire/models"
	uuid "github.com/satori/go.uuid"
)

func main() {
	u := &models.User{
		ID:       uuid.NewV4(),
		Name:     "gregory tandiono",
		Username: "gtandiono",
		Type:     "admin",
		Password: "password",
	}

	models.NewUser(u)
	fmt.Printf("great success")
}
