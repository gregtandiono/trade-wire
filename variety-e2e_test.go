package main

import (
	"testing"

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
			"id": uuid.NewV4().String(),
		})
}
