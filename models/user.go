package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = viper.GetString("secretKey")

type User struct {
	ID          rune   `sql:"id" json:"id"`
	Username    string `sql:"username" json:"username"`
	Email       string `sql:"email" json:"email"`
	Password    string `sql:"password"  json:"-"`
	RealName    string `sql:"real_name" json:"real_name"`
	City        string `sql:"city" json:"city"`
	Age         rune   `sql:"age" json:"age"`
	GamePoints  rune   `sql:"game_points" json:"game_points"`
	IsDisableAd bool   `sql:"is_disable_ad" json:"is_disable_ad"`
}

func (u *User) GenToken() string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwtToken.SignedString([]byte(secretKey))
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
