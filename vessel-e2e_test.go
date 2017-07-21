package main

import (
	"testing"
	"trade-wire/adaptors"

	"github.com/kataras/iris/httptest"
	uuid "github.com/satori/go.uuid"
)

func TestVesselHandler(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/vessels").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":     uuid.NewV4().String(),
			"name":   "MV Gisela Oldendorf",
			"beam":   "21.5",
			"loa":    "20",
			"draft":  "12m",
			"status": "loading at loadport",
		}).Expect().Status(200).JSON().Equal(adaptors.ResponseTemplate("insert:success"))
}
