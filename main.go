package main

import (
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	controller "bitbucket.com/gregtandiono_/trade-wire/controllers"
)

func main() {
	db := adaptors.DBConnector()

	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(customLogger)

	app.Get("/login", controller.NewUserController(db).Login)
	// app.Get("/login", controller.NewUserController(db).Login)
	// app.Get("/login", func(ctx *iris.Context) {
	// 	ctx.HTML(iris.StatusForbidden, "<h1> Please click <a href='/debug/pprof'>here</a>")
	// })
	// app.Post("/register", controller.NewUserController(db).Register)

	app.Listen(":8080")

}
