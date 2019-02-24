package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	gameHttpDeliver "quizChallenge/game/delivery/http"
	gameRepo "quizChallenge/game/repository"
	gameUcase "quizChallenge/game/usecase"
	"quizChallenge/question"
	questionHttpDeliver "quizChallenge/question/delivery/http"
	questionRepo "quizChallenge/question/repository"
	questionUcase "quizChallenge/question/usecase"
	"quizChallenge/user"
	userHttpDeliver "quizChallenge/user/delivery/http"
	userRepo "quizChallenge/user/repository"
	userUcase "quizChallenge/user/usecase"
	"time"
)

func init() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "-config=<file>")
	flag.Parse()
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func main() {
	var userRepository user.Repository
	var questionRepository question.Repository
	if viper.GetBool(`database.use`) {
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
		userRepository = userRepo.NewPsqlUserRepository(db)
		questionRepository = questionRepo.NewPsqlQuestionRepository(db)
	}
	var dbMongo *mgo.Database
	if viper.GetBool(`mongoDb.use`) {
		sessionMongo, err := mgo.Dial(viper.GetString("mongoDb.host"))
		if err != nil {
			os.Exit(1)
		}
		dbMongo = sessionMongo.DB(viper.GetString("mongoDb.name"))
		userRepository = userRepo.NewMongoUserRepository(dbMongo)
		questionRepository = questionRepo.NewMongoQuestionRepository(dbMongo)
		defer sessionMongo.Close()
	}
	gameRepository := gameRepo.NewMongoGameRepository(dbMongo)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.Recover())

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	uu := userUcase.NewUserUsecase(userRepository, timeoutContext)
	qu := questionUcase.NewQuestionUsecase(questionRepository, timeoutContext)
	gu := gameUcase.NewGameUsecase(gameRepository, userRepository, timeoutContext)

	userHttpDeliver.NewUserHttpHandler(e, uu)
	questionHttpDeliver.NewQuestionHttpHandler(e, qu)
	gameHttpDeliver.NewGameHttpHandelr(e, gu)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
