package controller

import (
	"github.com/jinzhu/gorm"
)

type UserController struct {
	database *gorm.DB
}

func NewUserController(database *gorm.DB) *UserController {
	return &UserController{
		database: database,
	}
}

func (uc *UserController) Register() {

}
