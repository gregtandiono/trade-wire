package main

import (
	"os"
	"testing"

	"trade-wire/fixtures"

	"gopkg.in/kataras/iris.v6"
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
	authResObj := e.POST("/auth").WithJSON(map[string]string{
		"username": "ghandsometon",
		"password": "password123",
	}).Expect().Status(200).JSON().Object()

	authResObj.Keys().ContainsOnly("id", "token")

	// It should fail on a bad login credential
	e.POST("/auth").WithJSON(u["invalidUserLogin"]).Expect().Status(400).JSON().Equal(map[string]string{
		"error": "username and password do not match",
	})

	// A user should be able to fetch their own user info
	aro := fetchToken(app, t)
	e.GET("/users/me").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"id":       aro["id"],
		"name":     "George Handsometon",
		"username": "ghandsometon",
		"type":     "admin",
		"password": "",
	})

	// A user should be able to update their own data and should return the updated field only
	e.PUT("/users/"+aro["id"]).
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "George Handsometon the second",
		}).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"id":       aro["id"],
		"name":     "George Handsometon the second",
		"username": "",
		"type":     "",
		"password": "",
	})

	// A user SHOULD only be able to update their own data, and server should validate it properly
	e.PUT("/users/569f9e19-5e3c-420b-8c7d-874529b50551").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "George Handsometon the second",
		}).
		Expect().
		Status(400).JSON().Equal(map[string]string{
		"error": "cannot update other users",
	})

}

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
