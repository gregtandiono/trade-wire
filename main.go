package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	// iris "gopkg.in/kataras/iris.v6"
	// "gopkg.in/kataras/iris.v6/adaptors/httprouter"
	// "gopkg.in/kataras/iris.v6/middleware/logger"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"

	"trade-wire/adaptors"

	controller "trade-wire/controllers"
)

// IrisHandler returns an instance of an iris framework
// baked into the main package so we can test the endpoints properly
func irisHandler() *iris.Application {
	app := iris.New()
	// app.Adapt(iris.DevLogger())
	// app.Adapt(httprouter.New())

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

	companies := app.Party("/companies", myJwtMiddleware.Serve)
	{
		companies.Post("", controller.NewCompanyController().Save)
		companies.Get("", controller.NewCompanyController().FetchAll)
		companies.Get("/:id", controller.NewCompanyController().FetchOne)
		companies.Put("/:id", controller.NewCompanyController().Update)
		companies.Delete("/:id", controller.NewCompanyController().Delete)
	}

	contacts := app.Party("/contacts", myJwtMiddleware.Serve)
	{
		contacts.Post("", controller.NewContactController().Save)
		contacts.Get("", controller.NewContactController().FetchAll)
		contacts.Get("/:id", controller.NewContactController().FetchOne)
		contacts.Put("/:id", controller.NewContactController().Update)
		contacts.Delete("/:id", controller.NewContactController().Delete)
	}

	commodities := app.Party("/commodities", myJwtMiddleware.Serve)
	{
		commodities.Post("", controller.NewCommodityController().Save)
		commodities.Get("", controller.NewCommodityController().FetchAll)
		commodities.Get("/:id", controller.NewCommodityController().FetchOne)
		commodities.Put("/:id", controller.NewCommodityController().Update)
		commodities.Delete("/:id", controller.NewCommodityController().Delete)
	}

	varieties := app.Party("/varieties", myJwtMiddleware.Serve)
	{
		varieties.Post("", controller.NewVarietyController().Save)
		varieties.Get("", controller.NewVarietyController().FetchAll)
		varieties.Get("/:id", controller.NewVarietyController().FetchOne)
		varieties.Put("/:id", controller.NewVarietyController().Update)
		varieties.Delete("/:id", controller.NewVarietyController().Delete)
	}

	return app
}

func main() {

	app := irisHandler()
	port, _, _ := adaptors.GetEnvironmentVariables()

	app.Listen(":" + port)

}
