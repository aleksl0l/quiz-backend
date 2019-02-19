package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
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
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "-config <file>")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	sourceName := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable",
		dbName,
		dbUser,
		dbPass,
		dbHost,
	)
	db, err := sql.Open("postgres", sourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	userRepository := userRepo.NewPsqlUserRepository(db)
	questionRepository := questionRepo.NewPsqlQuestionRepository(db)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	uu := userUcase.NewUserUsecase(userRepository, timeoutContext)
	qu := questionUcase.NewQuestionUsecase(questionRepository, timeoutContext)
	userHttpDeliver.NewUserHttpHandler(e, uu)
	questionHttpDeliver.NewQuestionHttpHandler(e, qu)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}

