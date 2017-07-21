package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
	uuid "github.com/satori/go.uuid"
)

func TestCompanyHandler(t *testing.T) {
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

	e.POST("/companies").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"id":           uuid.NewV4().String(),
			"name":         "charoen pokphand",
			"address":      "muara karang blok L9B no 12",
			"company_type": "buyer",
		}).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "company successfully created",
	})

	// A user should be able to fetch all companies
	e.GET("/companies").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Array().Length().
		Equal(33)

	// A user should be able to fetch a buyer
	companyObj := e.GET("/companies/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Object()

	companyObj.Value("name").Equal("Japfa Comfeed Indonesia")

	// A user should be able to update an existing buyer record
	companyUpdatedRecord := e.PUT("/companies/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		WithJSON(map[string]string{
			"name": "Japfa Comfeed Indonesia Tbk.",
		}).
		Expect().
		Status(200).JSON().Object()

	companyUpdatedRecord.Value("name").Equal("Japfa Comfeed Indonesia Tbk.")

	// A user should be able to soft delete a buyer record
	e.DELETE("/companies/f40e4dd4-f441-428b-8ff3-f893cb176819").
		WithHeader("Authorization", "Bearer "+aro["token"]).
		Expect().
		Status(200).JSON().Equal(map[string]string{
		"message": "company record successfully deleted",
	})
}
