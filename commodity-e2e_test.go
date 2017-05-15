package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestCommodityHandler(t *testing.T) {
	seedDataBase(t)
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/commodities").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":   uuid.NewV4().String(),
			"name": "Soybean",
		}).
		Expect().Status(200).JSON().Equal(map[string]string{
		"message": "commodity successfully created",
	})
}
