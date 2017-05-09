package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	iris "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/middleware/logger"

	"trade-wire/adaptors"

	controller "trade-wire/controllers"
)

// IrisHandler returns an instance of an iris framework
// baked into the main package so we can test the endpoints properly
func irisHandler() *iris.Framework {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	_, hashString, _ := adaptors.GetEnvironmentVariables()

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
	// WIP token validation
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
		users.Get("/me", controller.NewUserController().Me)
		users.Put("/:id", controller.NewUserController().Update)
		users.Delete("/:id", controller.NewUserController().Delete)
	}

	buyers := app.Party("/buyers", myJwtMiddleware.Serve)
	{
		buyers.Post("", controller.NewBuyerController().Save)
		buyers.Get("", controller.NewBuyerController().FetchAll)
		buyers.Get("/:id", controller.NewBuyerController().FetchOne)
		buyers.Put("/:id", controller.NewBuyerController().Update)
		buyers.Delete("/:id", controller.NewBuyerController().Delete)
	}

	return app
}

func main() {

	app := irisHandler()
	port, _, _ := adaptors.GetEnvironmentVariables()

	app.Listen(":" + port)

}
