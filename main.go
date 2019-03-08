package main

import (
	"flag"
	"github.com/01walid/echosentry"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"os"
	gameHttpDeliver "quizChallenge/game/delivery/http"
	gameRepo "quizChallenge/game/repository"
	gameUcase "quizChallenge/game/usecase"
	questionHttpDeliver "quizChallenge/question/delivery/http"
	questionRepo "quizChallenge/question/repository"
	questionUcase "quizChallenge/question/usecase"
	userHttpDeliver "quizChallenge/user/delivery/http"
	userRepo "quizChallenge/user/repository"
	userUcase "quizChallenge/user/usecase"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.json", "-config=<file>")
	flag.Parse()
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	var dbMongo *mgo.Database
	sessionMongo, err := mgo.Dial(viper.GetString("mongoDb.host"))
	if err != nil {
		os.Exit(1)
	}
	dbMongo = sessionMongo.DB(viper.GetString("mongoDb.name"))
	defer sessionMongo.Close()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
		echosentry.SetDSN(viper.GetString(`sentry.dsn`))
	e.Use(echosentry.Middleware())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${path} ${latency_human}` + "\n",
	}))

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepository := userRepo.NewMongoUserRepository(dbMongo)
	questionRepository := questionRepo.NewMongoQuestionRepository(dbMongo)
	gameRepository := gameRepo.NewMongoGameRepository(dbMongo)
	uu := userUcase.NewUserUsecase(userRepository, timeoutContext)
	qu := questionUcase.NewQuestionUsecase(questionRepository, timeoutContext)
	gu := gameUcase.NewGameUsecase(gameRepository, userRepository, questionRepository, timeoutContext)

	userHttpDeliver.NewUserHttpHandler(e, uu)
	questionHttpDeliver.NewQuestionHttpHandler(e, qu)
	gameHttpDeliver.NewGameHttpHandelr(e, gu)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))
}
