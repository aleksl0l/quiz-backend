package usecase

import (
	"context"
	"quizChallenge/game"
	"quizChallenge/models"
	"quizChallenge/user"
	"time"
)

type gameUsecase struct {
	gameRepo game.Repository
	userRepo user.Repository
	contextTimeout time.Duration
}

func NewGameUsecase(gr game.Repository, ur user.Repository, timeout time.Duration) game.Usecase {
	return &gameUsecase{
		gameRepo: gr,
		userRepo: ur,
		contextTimeout: timeout,
	}
}

func (u *gameUsecase) SearchGame(ctx context.Context, userId, gameType, gameCategory string) (*models.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.gameRepo.SearchGame(ctx, userId, gameType, gameCategory)
}

func (u *gameUsecase) GetGames(ctx context.Context, userId string) ([]*models.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.gameRepo.GetGames(ctx, userId)
}