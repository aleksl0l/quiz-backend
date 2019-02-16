package http

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"quizChallenge/models"
	"quizChallenge/user"
)

type HttpUserHandler struct {
	UUsecase user.Usecase
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RealName string `json:"real_name"`
	City     string `json:"city"`
	Age      rune   `json:"age"`
}

type JSONResponse map[string]interface{}

func NewUserHttpHandler(e *echo.Echo, us user.Usecase) {
	handler := &HttpUserHandler{
		UUsecase: us,
	}
	e.POST("/users/login", handler.Login)
	e.POST("/users", handler.SignUp)
}

func (u *HttpUserHandler) Login(c echo.Context) error {
	request := new(loginRequest)
	if err := c.Bind(request); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	uGet, _ := u.UUsecase.GetByUsername(ctx, request.Username)
	if err := uGet.CheckPassword(request.Password); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	response := JSONResponse{"token": uGet.GenToken()}
	response["user"] = uGet
	return c.JSON(http.StatusOK, response)
}

func (u *HttpUserHandler) SignUp(c echo.Context) error {
	signUpRequest := new(signUpRequest)
	if err := c.Bind(signUpRequest); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	userInstance := &models.User{
		Username: signUpRequest.Username,
		Email:    signUpRequest.Email,
		RealName: signUpRequest.RealName,
		City:     signUpRequest.City,
		Age:      signUpRequest.Age,
	}
	userInstance.SetPassword(signUpRequest.Password)
	if err := u.UUsecase.Store(ctx, userInstance); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User with that email or username is exist")
	}
	return c.JSON(http.StatusCreated, JSONResponse{"token": userInstance.GenToken()})
}
