package user

import (
	"context"
	"quizChallenge/models"
)

type Repository interface {
	//Fetch(ctx context.Context, cursor string, num int64) (res []*models.Article, nextCursor string, err error)
	//GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	//Update(ctx context.Context, ar *models.User) error
	Store(ctx context.Context, a *models.User) error
}
