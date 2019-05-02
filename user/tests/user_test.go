package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"quizChallenge/models"
	"quizChallenge/tests"
	"quizChallenge/user/repository"
	"quizChallenge/user/usecase"
	"testing"
	"time"
)

func GetTestUser() *models.User {
	return &models.User{
		Username:    "test_username",
		Email:       "test@quiz.com",
		Password:    "StrongPassword",
		RealName:    "Realname",
		City:        "NN",
		Age:         23,
		GamePoints:  10,
		IsDisableAd: false,
	}
}

func TestUserRepository(t *testing.T) {
	session, db := tests.GetTestDB()
	defer tests.FreeTestDB(session, db)

	asserts := assert.New(t)

	userRepository := repository.NewMongoUserRepository(db)
	user := GetTestUser()
	ctx := context.Background()
	err := userRepository.Store(ctx, user)
	asserts.Equal(nil, err, "Store should return nil error")

	var userGet *models.User
	userGet, err = userRepository.GetByUsername(ctx, user.Username)
	asserts.Equal(nil, err, "Getting should return nil error")
	asserts.Equal(userGet.Username, user.Username, "Username should equals")
}

func TestUserUsecase(t *testing.T) {
	session, db := tests.GetTestDB()
	defer tests.FreeTestDB(session, db)

	asserts := assert.New(t)
	timeout := time.Second * 2
	userRepository := repository.NewMongoUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, timeout)

	user := GetTestUser()
	ctx := context.Background()
	err := userUsecase.Store(ctx, user)
	asserts.Equal(nil, err, "Store should return nil error")

	var userGet *models.User
	userGet, err = userRepository.GetByUsername(ctx, user.Username)
	asserts.Equal(nil, err, "Getting should return nil error")
	asserts.Equal(userGet.Username, user.Username, "Username should equals")
}
