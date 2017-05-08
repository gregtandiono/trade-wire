package models

import (
	"errors"
	"time"
	"trade-wire/adaptors"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Type     string    `json:"type"`
	Password string    `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

func NewUserLogin(username, password string) *UserLogin {
	return &UserLogin{
		Username: username,
		Password: password,
	}
}

// HashPassword hashes password field from incoming requests
func (u *User) hashPassword() ([]byte, error) {
	hfp, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	return hfp, err
}

// Save saves new user instance into the DB
func (u *User) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	p, hperr := u.hashPassword()
	if hperr != nil {
		return hperr
	}

	uerr := u.checkIfUsernameExists()
	if uerr != nil {
		return uerr
	}
	db.Table("users").Create(&User{
		u.ID,
		u.Name,
		u.Username,
		u.Type,
		string(p),
	})

	return nil
}

func FetchAllUsers() []User {
	db := adaptors.DBConnector()
	defer db.Close()

	var users []User
	db.Select([]string{"id", "name", "username", "type"}).Where("deleted_at is null").Find(&users)
	return users
}

// FetchOne model method to fetch one user
// returns a map of one user
func (u *User) FetchOne() User {
	db := adaptors.DBConnector()
	defer db.Close()

	var user User
	db.Select([]string{"id", "name", "username", "type"}).Where("id = ?", u.ID).Find(&user)
	return user
}

// Update model method updates one user record
func (u *User) Update() {
	db := adaptors.DBConnector()
	defer db.Close()
	db.Table("users").Where("id = ?", u.ID).Updates(&u)
}

// Delete model method soft deletes user record
// it inserts a timestamp into the deleted_at column
func (u *User) Delete() {
	db := adaptors.DBConnector()
	defer db.Close()
	db.Table("users").Where("id = ?", u.ID).Update("deleted_at", time.Now())
}

func (ul *UserLogin) Auth() (map[string]string, error) {

	r, err := ul.checkPasswordAndGenerateTokenObject()
	return r, err
}

func (ul *UserLogin) generateToken(id uuid.UUID) string {
	_, hashString, _ := adaptors.GetEnvironmentVariables()
	var mySigningKey = []byte(hashString)
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

func (ul *UserLogin) checkPasswordAndGenerateTokenObject() (map[string]string, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	user, err := ul.checkForUser()
	if err != nil {
		return map[string]string{}, err
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ul.Password))
	if passErr != nil {
		return map[string]string{}, passErr
	}

	token := ul.generateToken(user.ID)

	r := map[string]string{
		"id":    uuid.UUID.String(user.ID),
		"token": token,
	}

	return r, nil
}

func (ul *UserLogin) checkForUser() (*User, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var user User
	var err error
	db.Table("users").Where("username = ?", ul.Username).Find(&user)
	if user.Name == "" {
		err = errors.New("user not found")
	}
	return &user, err
}

func (u *User) checkIfUsernameExists() error {
	db := adaptors.DBConnector()
	defer db.Close()

	var user User
	db.Table("users").Where("username = ?", u.Username).Find(&user)
	if user.Name != "" {
		return errors.New("user already exists")
	}

	return nil
}
