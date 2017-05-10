package main

import (
	"testing"

	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/httptest"
)

func fetchToken(app *iris.Framework, t *testing.T) map[string]string {
	e := httptest.New(app, t)

	// Assign variable to json response

	authResObj := e.POST("/auth").WithJSON(map[string]string{
		"username": "ghandsometon",
		"password": "password123",
	}).Expect().Status(200).JSON().Object()

	return map[string]string{
		"id":    authResObj.Value("id").String().Raw(),
		"token": authResObj.Value("token").String().Raw(),
	}

}
