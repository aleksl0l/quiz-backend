package http

import (
	"context"
	"github.com/labstack/echo"
	"net/http"
	"quizChallenge/models"
	"quizChallenge/question"
)

type HttpQuestionHandler struct {
	QUsecase question.Usecase
}

type questionGetRequest struct {
	Type     string `query:"type"`
	Category string `query:"category"`
}

func NewQuestionHttpHandler(e *echo.Echo, qu question.Usecase) {
	handler := &HttpQuestionHandler{
		QUsecase: qu,
	}
	e.GET("/questions", handler.GetQuestions)
	e.POST("/questions", handler.Store)
}

func (u *HttpQuestionHandler) GetQuestions(c echo.Context) error {
	request := &questionGetRequest{}
	if err := c.Bind(request); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	questions, err := u.QUsecase.GetQuestions(ctx, request.Type, request.Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error while fetching")
	}
	return c.JSON(http.StatusOK, questions)
}

func (u *HttpQuestionHandler) Store(c echo.Context) error {
	request := &models.Question{}
	if err := c.Bind(request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err := u.QUsecase.Store(ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error while fetching")
	}
	return c.JSON(http.StatusOK, request)
}