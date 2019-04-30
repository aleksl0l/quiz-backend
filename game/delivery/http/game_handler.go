package http

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"net/http"
	"quizChallenge/game"
)

type HttpGameHandler struct {
	GUsecase game.Usecase
}

type searchGameRequest struct {
	Type     string `json:"type"`
	Category string `json:"category"`
}

func NewGameHttpHandler(e *echo.Echo, gu game.Usecase) {
	handler := &HttpGameHandler{
		GUsecase: gu,
	}
	gameGroup := e.Group("/games")
	gameGroup.Use(middleware.JWT([]byte(viper.GetString(`secretKey`))))
	gameGroup.POST("/search_game", handler.SearchGame)
	gameGroup.GET("", handler.GetGames)
	gameGroup.GET("/:gameId", handler.GetGameById)
}

func (u *HttpGameHandler) SearchGame(c echo.Context) error {
	request := &searchGameRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	game, err := u.GUsecase.SearchGame(ctx, userId, request.Type, request.Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, game)
}

func (u *HttpGameHandler) GetGames(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	games, err := u.GUsecase.GetGames(ctx, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, games)
}

func (u *HttpGameHandler) GetGameById(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	gameId := c.Param("gameId")
	game, err := u.GUsecase.GetGameById(ctx, gameId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, game)
}
