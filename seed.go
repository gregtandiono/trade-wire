package main

import (
	"fmt"
	"testing"
	"trade-wire/adaptors"

	randomdata "github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"

	"trade-wire/fixtures"

	"github.com/kataras/iris/httptest"
)

func seedDataBase(t *testing.T) {
	destroyTables()
	seedUsers(t)
	seedCompanies(t)
	seedContacts(t)
	seedCommodities(t)
	seedVarieties(t)
}

func destroyTables() {
	db := adaptors.DBConnector()
	defer db.Close()

	tables := []string{
		"trades",
		"vessels",
		"contacts",
		"companies",
		"varieties",
		"commodities",
		"companies",
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

func seedCompanies(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	b := fixtures.BuyerFixtures()

	// Seed Known buyer
	au := fetchToken(app, t)

	e.POST("/companies").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(b["validBuyerRecord"]).
		Expect().Status(200)

	for i := 0; i < 20; i++ {
		e.POST("/companies").
			WithHeader("Authorization", "Bearer "+au["token"]).
			WithJSON(map[string]string{
				"id":           uuid.NewV4().String(),
				"name":         randomdata.SillyName(),
				"address":      randomdata.Address(),
				"company_type": "buyer",
			}).Expect().Status(200)
	}

	for i := 0; i < 10; i++ {
		e.POST("/companies").
			WithHeader("Authorization", "Bearer "+au["token"]).
			WithJSON(map[string]string{
				"id":           uuid.NewV4().String(),
				"name":         randomdata.SillyName(),
				"address":      randomdata.Address(),
				"company_type": "supplier",
			}).Expect().Status(200)
	}
}

func seedCommodities(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	au := fetchToken(app, t)
	c := fixtures.CommodityFixtures()

	e.POST("/commodities").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(c["validCommodityRecord"]).Expect().Status(200)

	e.POST("/commodities").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(c["validCommodityRecord2"]).Expect().Status(200)
}

func seedVarieties(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	au := fetchToken(app, t)
	v := fixtures.VarietyFixtures()

	e.POST("/varieties").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(v["validVarietyRecord"]).Expect().Status(200)

	e.POST("/varieties").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(v["validVarietyRecord2"]).Expect().Status(200)
}

func seedContacts(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	au := fetchToken(app, t)

	e.POST("/contacts").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(map[string]string{
			"id":            "4ce32ff4-7fe3-49b9-b40b-d4b3a782696d",
			"name":          "Dewi Tjandra",
			"position":      "Head of Operations",
			"office_number": "+622178291827",
			"cell_number":   "+628117920029",
			"notes":         "lorem ipsum dolor sit amet",
			"company_id":    "f40e4dd4-f441-428b-8ff3-f893cb176819",
		}).Expect().Status(200)

	for i := 0; i < 5; i++ {
		e.POST("/contacts").
			WithHeader("Authorization", "Bearer "+au["token"]).
			WithJSON(map[string]string{
				"id":            uuid.NewV4().String(),
				"name":          randomdata.SillyName(),
				"position":      randomdata.Adjective(),
				"office_number": "+622178291827",
				"cell_number":   "+628117920029",
				"notes":         "lorem ipsum dolor sit amet",
				"company_id":    "f40e4dd4-f441-428b-8ff3-f893cb176819",
			}).Expect().Status(200)

	}
}

func seedVessels(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	au := fetchToken(app, t)

	e.POST("/vessels").
		WithHeader("Authorization", "Bearer "+au["token"]).
		WithJSON(map[string]string{
			"id":     "f39165b1-15ba-412e-822c-419d508a7348",
			"name":   "MV Marlin",
			"beam":   "22.4",
			"loa":    "18",
			"draft":  "11m",
			"status": "berth at loading port",
		}).Expect().Status(200)

	for i := 0; i < 5; i++ {
		e.POST("/vessels").
			WithHeader("Authorization", "Bearer "+au["token"]).
			WithJSON(map[string]string{
				"id":     uuid.NewV4().String(),
				"name":   randomdata.SillyName(),
				"beam":   "22.4",
				"loa":    "18",
				"draft":  "11",
				"status": "berth at loading port",
			}).Expect().Status(200)
	}
}

func seedTrades(t *testing.T) {
}
