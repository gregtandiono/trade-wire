package controller

import (
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"

	"trade-wire/models"
)

// UserController struct serves as a initializer
type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Login(ctx context.Context) {
	// var userLogin models.UserLogin
	userLogin := &models.UserLogin{}
	ctx.ReadJSON(userLogin)
	ul := models.NewUserLogin(userLogin.Username, userLogin.Password)
	tokenObj, err := ul.Auth()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "username and password do not match",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(&tokenObj)
	}
}

func (uc *UserController) FetchAll(ctx context.Context) {
	users := models.FetchAllUsers()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(users)
}

func (uc *UserController) Me(ctx context.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	ur, err := user.Me(fetchTokenFromHeader(ctx))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"message": err.Error(),
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(ur)
	}
}

func (uc *UserController) Update(ctx context.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Params().Get("id")
	user.ID = uuid.FromStringOrNil(userID)
	err := user.Update(fetchTokenFromHeader(ctx))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(&user)
	}
}

func (uc *UserController) Delete(ctx context.Context) {
	var user models.User
	ctx.ReadJSON(&user)
	userID := ctx.Params().Get("id")
	user.ID = uuid.FromStringOrNil(userID)
	err := user.Delete(fetchTokenFromHeader(ctx))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "user successfully deleted",
		})
	}
}

func (uc *UserController) Register(ctx context.Context) {
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
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "user successfully registered",
		})
	}
}

func fetchTokenFromHeader(ctx context.Context) (tokenString string) {
	header := strings.Split(ctx.GetHeader("Authorization"), " ")
	return header[1]
}
