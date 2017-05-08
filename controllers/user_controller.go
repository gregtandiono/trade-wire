package controller

import (
	uuid "github.com/satori/go.uuid"

	"trade-wire/models"

	"gopkg.in/kataras/iris.v6"

	"fmt"
)

// UserController struct serves as a initializer
type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Login(ctx *iris.Context) {
	// var userLogin models.UserLogin
	userLogin := &models.UserLogin{}
	ctx.ReadJSON(userLogin)
	ul := models.NewUserLogin(userLogin.Username, userLogin.Password)
	tokenObj, err := ul.Auth()
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "username and password do not match",
		})
	} else {
		ctx.JSON(iris.StatusOK, &tokenObj)
	}
}

func (uc *UserController) FetchAll(ctx *iris.Context) {
	users := models.FetchAllUsers()
	ctx.JSON(iris.StatusOK, users)
}

func (uc *UserController) FetchOne(ctx *iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Param("id")
	user.ID = uuid.FromStringOrNil(userID)
	ur := user.FetchOne()
	ctx.JSON(iris.StatusOK, ur)
}

func (uc *UserController) Update(ctx *iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Param("id")
	user.ID = uuid.FromStringOrNil(userID)
	user.Update()
	ctx.JSON(iris.StatusOK, &user)
}

func (uc *UserController) Delete(ctx *iris.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Param("id")
	user.ID = uuid.FromStringOrNil(userID)
	user.Delete()
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
		user.Password,
	)
	err := u.Save()
	if err != nil {
		fmt.Print(err)
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(iris.StatusOK, map[string]string{
			"message": "user successfully registered",
		})
	}
}
