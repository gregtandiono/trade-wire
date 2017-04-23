package models

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"bitbucket.com/gregtandiono_/trade-wire/adaptors"
)

// User model
type User struct {
	ID                             uuid.UUID
	Name, Username, Type, Password string
}

// NewUser {u} is an instance of user struct
func NewUser(u *User) {

	db := adaptors.DBConnector()
	rows, err := db.Query(`
		INSERT INTO users (id, name, username, type, password)
		VALUES ($1, $2, $3, $4, $5)
	`, u.ID, u.Name, u.Username, u.Type, u.Password)

	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
}
