package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Password []byte    `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

// NewUser {u} is an instance of user struct
func NewUser(id uuid.UUID, name, username, t string, password []byte) *User {
	return &User{
		ID:       id,
		Name:     name,
		Username: username,
		Type:     t,
		Password: password,
	}
}

func NewUserLogin(username string, password []byte) *UserLogin {
	return &UserLogin{
		Username: username,
		Password: password,
	}
}

// HashPassword hashes password field from incoming requests
func (u *User) hashPassword() []byte {
	hfp, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hfp
}

func (u *User) Save(db *gorm.DB) {
	p := u.hashPassword()
	db.Table("users").Create(&User{
		u.ID,
		u.Name,
		u.Username,
		u.Type,
		p,
	})
}

func FetchAllUsers(db *gorm.DB) []User {
	var users []User
	db.Select([]string{"id", "name", "username", "type"}).Where("deleted_at is null").Find(&users)
	return users
}

func (u *User) Update(db *gorm.DB) {
	db.Table("users").Where("id = ?", u.ID).Updates(&u)
}

func (u *User) Delete(db *gorm.DB) {
	db.Table("users").Where("id = ?", u.ID).Update("deleted_at", time.Now())
}

func (ul *UserLogin) Auth(db *gorm.DB) map[string]string {
	return ul.checkPasswordAndGenerateTokenObject(db)
}

func (ul *UserLogin) generateToken(id uuid.UUID) string {
	var mySigningKey = []byte("supersecretkey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"id":  id,
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (ul *UserLogin) checkPasswordAndGenerateTokenObject(db *gorm.DB) map[string]string {
	user, err := ul.checkForUser(db)
	if err != "" {
		return map[string]string{
			"error": err,
		}
	}
	passErr := bcrypt.CompareHashAndPassword(user.Password, ul.Password)
	if passErr != nil {
		return map[string]string{
			"error": "password does not match",
		}
	}

	token := ul.generateToken(user.ID)

	return map[string]string{
		"id":    uuid.UUID.String(user.ID),
		"token": token,
	}
}

func (ul *UserLogin) checkForUser(db *gorm.DB) (*User, string) {
	// var data UserLogin
	var user User
	var err string
	db.Table("users").Where("username = ?", ul.Username).Find(&user)
	if user.Name == "" {
		err = "user not found"
	}
	return &user, err
}
