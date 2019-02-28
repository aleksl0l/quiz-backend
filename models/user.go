package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = viper.GetString("secretKey")

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"-"`
	RealName    string        `json:"real_name" bson:"realName"`
	City        string        `json:"city"`
	Age         rune          `json:"age"`
	GamePoints  rune          `json:"game_points" bson:"gamePoints"`
	IsDisableAd bool          `json:"is_disable_ad" bson:"isDisableAd"`
}

func (u *User) GenToken() string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"id":  u.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	secretKey := viper.GetString("secretKey")
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
