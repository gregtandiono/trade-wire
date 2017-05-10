package main

import (
	"testing"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestCommodityHandler(t *testing.T) {
	seedDataBase(t)
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/commodities")
	e.GET("/commodities")
}
