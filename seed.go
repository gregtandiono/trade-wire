package main

import (
	"fmt"
	"testing"
	"trade-wire/adaptors"

	randomdata "github.com/Pallinder/go-randomdata"

	"trade-wire/fixtures"

	"gopkg.in/kataras/iris.v6/httptest"
)

func seedDataBase(t *testing.T) {
	destroyTables()
	seedUsers(t)
}

func destroyTables() {
	db := adaptors.DBConnector()
	defer db.Close()

	tables := []string{
		"tracking",
		"contacts",
		"trades",
		"suppliers",
		"varieties",
		"commodities",
		"buyers",
		"users",
	}

	for _, table := range tables {
		fmt.Println("#### Deleting from table " + table)
		db.Exec("DELETE FROM " + table)
	}
}

func seedUsers(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)

	u := fixtures.UserFixtures()

	// Seed known user

	e.POST("/register").WithJSON(u["validUserSignup"]).Expect().Status(200)
	e.POST("/register").WithJSON(u["validEmployeeSignup"]).Expect().Status(200)

	// Seed random user

	for i := 0; i < 10; i++ {
		e.POST("/register").WithJSON(map[string]string{
			"name":     randomdata.FullName(randomdata.RandomGender),
			"username": randomdata.SillyName(),
			"type":     "employee",
			"password": "07jjpasimyh",
		}).Expect().Status(200)
	}
}
