package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SECRET_KEY = "A String Very Very Very Strong!!@##$!@#$"

type User struct {
	ID       rune   `sql:"id" json:"id"`
	Username string `sql:"username" json:"username"`
	Email    string `sql:"email" json:"email"`
	Password string `sql:"password"  json:"-"`
}

func (u *User) GenToken() string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwtToken.SignedString([]byte(SECRET_KEY))
	return token
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, 2)
	u.Password = string(passwordHash)
	return nil
}
