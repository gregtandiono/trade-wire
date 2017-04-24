package controller

import (
	"bitbucket.com/gregtandiono_/trade-wire/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/kataras/iris.v6"
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

func (uc *UserController) Register(ctx *iris.Context) {
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

func (uc *UserController) Login(ctx *iris.Context) {

}
