package main

import (
	"testing"
	"trade-wire/fixtures"

	"github.com/kataras/iris/httptest"
	uuid "github.com/satori/go.uuid"
)

func TestVarietyHandler(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/varieties").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":           uuid.NewV4().String(),
			"commodity_id": fixtures.CommodityFixtures()["validCommodityRecord"]["id"],
			"name":         "Australian Standard White (ASW)",
			"origin":       "australia",
			"specs":        "Protein min. 10.50%\n Moisture max. 13.50%\n Foreign matter max. 1.00%\n Bug Damage max.1.00%\n Test Weight min. 76kg/HL\n Aflatoxin max. 20ppb\n Vomitoxin max.2ppm",
		}).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "variety successfully created",
	})
}
