package models

import (
	"golang.org/x/crypto/bcrypt"

	"database/sql"

	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID                             uuid.UUID
	Name, Username, Type, Password string
}

// NewUser {u} is an instance of user struct
func NewUser(id uuid.UUID, name, username, t, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Username: username,
		Type:     t,
		Password: password,
	}
}

func (u *User) HashPassword() []byte {
	password := []byte(u.Password)
	hfp, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hfp
}

func (u *User) Save(db *sql.DB) *sql.Rows {

	p := u.HashPassword()
	rows, err := db.Query(`
		INSERT INTO users (id, name, username, type, password)
		VALUES ($1, $2, $3, $4, $5)
	`, u.ID, u.Name, u.Username, u.Type, p)

	if err != nil {
		panic(err)
	}

	return rows
}

func (u *User) Update(db *sql.DB, id uuid.UUID, ud *User) {
}

func (u *User) Destroy(db *sql.DB) {
}
