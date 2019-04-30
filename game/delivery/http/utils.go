package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getUserIdFromContext(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	return userId
}