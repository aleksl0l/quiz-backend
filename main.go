package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	questionHttpDeliver "quizChallenge/question/delivery/http"
	questionRepo "quizChallenge/question/repository"
	questionUcase "quizChallenge/question/usecase"
	userHttpDeliver "quizChallenge/user/delivery/http"
	userRepo "quizChallenge/user/repository"
	userUcase "quizChallenge/user/usecase"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("postgres", "dbname=quiz user=quiz password=quiz host=localhost sslmode=disable")
}

func main() {
	defer db.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	userRepository := userRepo.NewPsqlUserRepository(db)
	questionRepository := questionRepo.NewPsqlQuestionRepository(db)

	timeoutContext := time.Duration(60) * time.Second
	uu := userUcase.NewUserUsecase(userRepository, timeoutContext)
	qu := questionUcase.NewQuestionUsecase(questionRepository, timeoutContext)
	userHttpDeliver.NewUserHttpHandler(e, uu)
	questionHttpDeliver.NewQuestionHttpHandler(e, qu)

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
