package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
	uuid "github.com/satori/go.uuid"
)

func TestContactHandler(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/contacts").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":            uuid.NewV4().String(),
			"name":          "Arief Widjaja",
			"position":      "Head of Procurement",
			"office_number": "(021)668918827",
			"cell_number":   "+62812789372",
			"notes":         "lorem ipsum dolor sit amet",
			"company_id":    "f40e4dd4-f441-428b-8ff3-f893cb176819",
		}).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "contact successfully created",
	})

	// A user should be able to fetch all contacts
	e.GET("/contacts").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Array().Length().
		Equal(32)

	// A user should be able to fetch a buyer
	companyObj := e.GET("/contacts/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Object()

	companyObj.Value("name").Equal("Japfa Comfeed Indonesia")

	// A user should be able to update an existing buyer record
	companyUpdatedRecord := e.PUT("/contacts/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "Japfa Comfeed Indonesia Tbk.",
		}).
		Expect().
		Status(200).JSON().Object()

	companyUpdatedRecord.Value("name").Equal("Japfa Comfeed Indonesia Tbk.")

	// A user should be able to soft delete a buyer record
	e.DELETE("/contacts/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "company record successfully deleted",
	})
}
