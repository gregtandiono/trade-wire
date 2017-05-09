package main

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/kataras/iris.v6/httptest"
)

func TestBuyerHandler(t *testing.T) {
	seedDataBase(t)
	app := irisHandler()
	e := httptest.New(app, t)
	aro := fetchToken(app, t)

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
