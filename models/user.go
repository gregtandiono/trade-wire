package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID                             uuid.UUID
	Name, Username, Type, Password string
}

// NewUser {u} is an instance of user struct
func NewUser(u *User) {

	db := adaptors.DBConnector()
	p := u.hash()
	rows, err := db.Query(`
		INSERT INTO users (id, name, username, type, password)
		VALUES ($1, $2, $3, $4, $5)
	`, u.ID, u.Name, u.Username, u.Type, p)

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
}

func (u *User) hash() []byte {
	password := []byte(u.Password)
	hfp, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hfp
}
