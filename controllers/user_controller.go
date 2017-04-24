package controller

import (
	"fmt"

	"github.com/jinzhu/gorm"

	uuid "github.com/satori/go.uuid"

	"bitbucket.com/gregtandiono_/trade-wire/models"
	"gopkg.in/kataras/iris.v6"
	_ "gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

// UserController struct serves as a initializer
type UserController struct {
	database *gorm.DB
}

func NewUserController(database *gorm.DB) *UserController {
	return &UserController{
		database: database,
	}
}

func (uc *UserController) Login(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, map[string]string{"name": "gregory"})
}

func (uc *UserController) FetchAll(ctx *iris.Context) {
	models.FetchAll(uc.database)
	ctx.JSON(iris.StatusOK, map[string]string{"status": "ok"})
}

func (uc *UserController) Register(ctx *iris.Context) {
	fmt.Printf("hello")
	var user models.User
	ctx.ReadJSON(&user)
	u := models.NewUser(
		uuid.NewV4(),
		user.Name,
		user.Username,
		user.Type,
		[]byte(user.Password),
	)
	u.Save(uc.database)
	ctx.JSON(iris.StatusOK, &user)
}
