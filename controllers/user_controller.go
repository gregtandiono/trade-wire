package controller

import (
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
	var userLogin models.UserLogin
	ctx.ReadJSON(&userLogin)
	ul := models.NewUserLogin(userLogin.Username, userLogin.Password)
	tokenObj := ul.Auth(uc.database)
	ctx.JSON(iris.StatusOK, &tokenObj)
}

func (uc *UserController) FetchAll(ctx *iris.Context) {
	users := models.FetchAllUsers(uc.database)
	ctx.JSON(iris.StatusOK, users)
}

func (uc *UserController) Update(ctx *iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Param("id")
	user.ID = uuid.FromStringOrNil(userID)
	user.Update(uc.database)
	ctx.JSON(iris.StatusOK, &user)
}

func (uc *UserController) Delete(ctx *iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Param("id")
	user.ID = uuid.FromStringOrNil(userID)
	user.Delete(uc.database)
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "record successfully deleted",
	})
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
