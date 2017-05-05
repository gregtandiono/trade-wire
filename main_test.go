package main

import (
	"os"
	"testing"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestIrisHandler(t *testing.T) {
	env := os.Getenv("ENV")
	if env == "TEST" {
		seedDataBase(t)
	}

	app := irisHandler()
	e := httptest.New(app, t)

	e.GET("/auth").Expect().Status(404)
}
