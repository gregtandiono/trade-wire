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
	vid := fixtures.VarietyFixtures()["validVarietyRecord"]["id"]
	cid := fixtures.CommodityFixtures()["validCommodityRecord"]["id"]

	e.POST("/varieties").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":           uuid.NewV4().String(),
			"commodity_id": cid,
			"name":         "Australian Standard White (ASW)",
			"origin":       "australia",
			"specs":        "Protein min. 10.50%\n Moisture max. 13.50%\n Foreign matter max. 1.00%\n Bug Damage max.1.00%\n Test Weight min. 76kg/HL\n Aflatoxin max. 20ppb\n Vomitoxin max.2ppm",
		}).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "variety successfully created",
	})

	e.POST("/varieties").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":     uuid.NewV4().String(),
			"name":   "Australian Standard White (ASW)",
			"origin": "australia",
		}).
		Expect().Status(400).JSON().Equal(map[string]string{
		"error": "failed to insert variety record",
	})

	varietyObj := e.GET("/varieties/"+vid).
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().Status(200).JSON().Object()

	varietyObj.Value("name").Equal("Black Sea Wheat (BS Wheat)")
	varietyObj.Value("origin").Equal("Black Sea")

	varietyUpdatedRecord := e.PUT("/varieties/"+vid).
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "Ukraine Milling Wheat (UMW)",
		}).
		Expect().Status(200).JSON().Object()

	varietyUpdatedRecord.Value("name").Equal("Ukraine Milling Wheat (UMW)")

	e.DELETE("/varieties/"+vid).
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "record successfully deleted",
	})
}
