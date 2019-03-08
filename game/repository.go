package game

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/models"
)

type Repository interface {
	//SearchGame(ctx context.Context, userId, gameType, gameCategory string) (*models.Game, error)
	GetGames(ctx context.Context, userId string) ([]*models.Game, error)
	GetGamesByCondition(ctx context.Context, condition bson.M) (*models.Game, error)
	Store(ctx context.Context, game *models.Game) error
	Update(ctx context.Context, game *models.Game, update bson.M) error
}
