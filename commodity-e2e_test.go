package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestCommodityHandler(t *testing.T) {
	// seedDataBase(t)
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/commodities").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":   uuid.NewV4().String(),
			"name": "soybean",
		}).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "commodity successfully created",
	})

	e.GET("/commodities").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().Status(200).JSON().Array().Length().Equal(3)

	commodityObj := e.GET("/commodities/75a5cdfe-ca69-4680-a903-af89eaaa4804").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().Status(200).JSON().Object()

	commodityObj.Value("name").Equal("wheat")

	commodityUpdatedRecord := e.PUT("/commodities/75a5cdfe-ca69-4680-a903-af89eaaa4804").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "soybean meal",
		}).
		Expect().Status(200).JSON().Object()

	commodityUpdatedRecord.Value("name").Equal("soybean meal")

	e.DELETE("/commodities/75a5cdfe-ca69-4680-a903-af89eaaa4804").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "record successfully deleted",
	})
}
