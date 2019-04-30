package game

import (
	"context"
	"quizChallenge/models"
)

type Usecase interface {
	SearchGame(ctx context.Context, userId, gameType, gameCategory string) (*models.Game, error)
	GetGames(ctx context.Context, userId string) ([]*models.Game, error)
	GetGameById(ctx context.Context, gameId string) (*models.Game, error)
	AnswerQuestion(ctx context.Context, gameId, questionId, userId, answer string) error
}
