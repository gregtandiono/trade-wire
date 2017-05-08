package main

import (
	"os"
	"testing"

	"trade-wire/fixtures"

	"gopkg.in/kataras/iris.v6/httptest"
)

func TestIrisHandler(t *testing.T) {
	env := os.Getenv("ENV")
	if env == "TEST" {
		seedDataBase(t)
	}

	app := irisHandler()
	e := httptest.New(app, t)
	u := fixtures.UserFixtures()

	// It should be able to register a new user
	e.POST("/register").WithJSON(map[string]string{
		"name":     "Robert Patrick",
		"username": "rpatrick",
		"type":     "employee",
		"password": "my037yhh",
	}).Expect().Status(200).JSON().Equal(map[string]string{
		"message": "user successfully registered",
	})

	// It should fail on registering a new user with the same username
	e.POST("/register").WithJSON(map[string]string{
		"name":     "Awesome Johnson",
		"username": "rpatrick",
		"type":     "employee",
		"password": "my037yhh",
	}).Expect().Status(400).JSON().Equal(map[string]string{
		"error": "user already exists",
	})

	// It should be able to login an existing admin user
	e.POST("/auth").WithJSON(map[string]string{
		"username": "ghandsometon",
		"password": "password123",
	}).Expect().Status(200).JSON()

	// It should fail on a bad login credential
	e.POST("/auth").WithJSON(u["invalidUserLogin"]).Expect().Status(400).JSON().Equal(map[string]string{
		"error": "username and password do not match",
	})
}
