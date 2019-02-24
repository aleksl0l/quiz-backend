package user

import (
	"context"
	"quizChallenge/models"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Store(ctx context.Context, a *models.User) error
}
