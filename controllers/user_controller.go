package controller

import (
	"github.com/jinzhu/gorm"
)

type UserController struct {
	database gorm.DB
}
