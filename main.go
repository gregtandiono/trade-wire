package main

import (
	controller "bitbucket.com/gregtandiono_/trade-wire/controllers"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
)

func main() {
	db := adaptors.DBConnector()
	// u := models.NewUser(
	// 	uuid.NewV4(),
	// 	"gregory tandiono",
	// 	"gtandiono",
	// 	"admin",
	// 	[]byte("password"),
	// )

	// u.Save(db)

	app := iris.New()
	app.Adapt(httprouter.New())

	app.Post("/register", controller.NewUserController(db).Register)

	app.Listen(":8080")

}
