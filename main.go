package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	controller "bitbucket.com/gregtandiono_/trade-wire/controllers"
)

func main() {

	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	port, hashString, _ := adaptors.GetEnvironmentVariables()

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

	// auth middleware auth
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(hashString), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app.Use(customLogger)

	// Pre-app endpoints
	app.Post("/auth", controller.NewUserController().Login)
	app.Post("/register", controller.NewUserController().Register)

	// user endpoints
	users := app.Party("/users", myJwtMiddleware.Serve)
	{
		users.Get("/", controller.NewUserController().FetchAll)
		users.Get("/:id", controller.NewUserController().FetchOne)
		users.Put("/:id", controller.NewUserController().Update)
		users.Delete("/:id", controller.NewUserController().Delete)
	}

	app.Listen(":" + port)

}
