package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/game"
	"quizChallenge/models"
	"quizChallenge/question"
	"quizChallenge/user"
	"time"
)

type gameUsecase struct {
	gameRepo       game.Repository
	userRepo       user.Repository
	questionRepo   question.Repository
	contextTimeout time.Duration
}

func NewGameUsecase(gr game.Repository, ur user.Repository, qr question.Repository, timeout time.Duration) game.Usecase {
	return &gameUsecase{
		gameRepo:       gr,
		userRepo:       ur,
		questionRepo:   qr,
		contextTimeout: timeout,
	}
}

func (u *gameUsecase) SearchGame(ctx context.Context, userId, gameType, gameCategory string) (*models.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	game := &models.Game{}
	game, err := u.gameRepo.GetGamesByCondition(ctx, bson.M{
		"typeQuestions":     gameType,
		"categoryQuestions": gameCategory,
		"user2id":           nil,
	})
	now := time.Now()
	if err != nil {
		game = &models.Game{}
		game.ID = bson.NewObjectId()
		game.TypeQuestions = gameType
		game.CategoryQuestions = gameCategory
		game.User1ID = bson.ObjectIdHex(userId)
		game.StartedAt = &now
		game.Questions, _ = u.questionRepo.GetRandomQuestions(ctx, game.TypeQuestions, game.CategoryQuestions, 7)
		game.UsersAnswers = make([]*models.UserAnswer, 0, 7)
		for _, q := range game.Questions {
			game.UsersAnswers = append(game.UsersAnswers, &models.UserAnswer{QuestionId: q.ID})
		}
		err := u.gameRepo.Store(ctx, game)
		if err != nil {
			return nil, err
		}
	} else {
		if game.User1ID.Hex() == userId {
			return nil, errors.New("you already search game")
		}
		err = u.gameRepo.Update(ctx, game, bson.M{"user2id": bson.ObjectIdHex(userId)})
		if err != nil {
			return nil, err
		}
		game.User2ID = bson.ObjectIdHex(userId)
	}
	return game, nil
}

func (u *gameUsecase) GetGames(ctx context.Context, userId string) ([]*models.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.gameRepo.GetGames(ctx, userId)
}

func (u *gameUsecase) GetGameById(ctx context.Context, gameId string) (*models.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	game, err := u.gameRepo.GetGamesByCondition(ctx, bson.M{"_id": bson.ObjectIdHex(gameId)})
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (u *gameUsecase) AnswerQuestion(ctx context.Context, gameId, questionId, userId, answer string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	game, err := u.gameRepo.GetGameByGameIdQuestionId(ctx, gameId, questionId)
	if err != nil {
		return err
	}
	var userNum int8
	if bson.ObjectIdHex(userId) == game.User1ID {
		userNum = 1
	} else {
		userNum = 2
	}
	questionObjectId := bson.ObjectIdHex(questionId)
	for i, val := range game.UsersAnswers {
		if val.QuestionId ==  questionObjectId {

			updatedField := fmt.Sprintf("users_answers.%d.user%d_answer", i, userNum)
			return u.gameRepo.Update(ctx, game, bson.M{updatedField: answer})
		}
	}
	return errors.New("not found")
}