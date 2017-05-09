package main

import (
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"

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

	// A user SHOULD be able to disable their account, WIP on a better process through admin
	e.DELETE("/users/"+aro["id"]).
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "user successfully deleted",
	})

	// A user SHOULD only be able to delete themseleves and nobody else, and server should validate that
	e.DELETE("/users/569f9e19-5e3c-420b-8c7d-874529b50551").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(400).JSON().Equal(map[string]string{
		"error": "cannot delete other users",
	})

	// A user should be able to create a buyer
	e.POST("/buyers").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":      uuid.NewV4().String(),
			"name":    "charoen pokphand",
			"address": "muara karang blok L9B no 12",
			"pic": `[
				{"name": "hendra tjang", "telephone": "6281237738777"},
				{"name": "felicia kurniawan", "telephone": "76632888"}
			]`,
		}).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "buyer successfully created",
	})

	// A user should be able to fetch all buyers
	e.GET("/buyers").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Array().Length().
		Equal(22)

	// A user should be able to fetch a buyer
	buyerObj := e.GET("/buyers/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Object()

	buyerObj.Value("name").Equal("Japfa Comfeed Indonesia")

	// A user should be able to update an existing buyer record
	buyerUpdatedRecord := e.PUT("/buyers/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "Japfa Comfeed Indonesia Tbk.",
		}).
		Expect().
		Status(200).JSON().Object()

	buyerUpdatedRecord.Value("name").Equal("Japfa Comfeed Indonesia Tbk.")

	// A user should be able to soft delete a buyer record
	e.DELETE("/buyers/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "buyer record successfully deleted",
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
