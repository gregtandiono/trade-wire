package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"database/sql"

	"fmt"

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

func (u *User) Authorize(db *gorm.DB) {
	u.checkForUser(db)
}

func (u *User) FetchAll(db *gorm.DB) *sql.Rows {
	rows, err := db.Find(u).Rows()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return rows
}

func (u *User) Update(db *sql.DB, id uuid.UUID, ud *User) {
}

func (u *User) Destroy(db *sql.DB) {
}

func (u *User) generateToken() string {
	var mySigningKey = []byte("supersecretkey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func (u *User) checkForUser(db *gorm.DB) {
	db.Find(&u, "72b0afdf-e283-41a7-9f65-540b0095b62c")
	fmt.Println(u.Name)
}
