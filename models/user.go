package models

import (
	"golang.org/x/crypto/bcrypt"

	"database/sql"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID                   uuid.UUID
	Name, Username, Type string
	Password             []byte
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

func (u *User) HashPassword() []byte {
	hfp, err := bcrypt.GenerateFromPassword(u.Password, bcrypt.DefaultCost)
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

func (u *User) SaveWithGorm(db *gorm.DB) {
	p := u.HashPassword()
	db.Table("users").Create(&User{
		u.ID,
		u.Name,
		u.Username,
		u.Type,
		p,
	})
	// defer db.Close()
}

func (u *User) Update(db *sql.DB, id uuid.UUID, ud *User) {
}

func (u *User) Destroy(db *sql.DB) {
}
