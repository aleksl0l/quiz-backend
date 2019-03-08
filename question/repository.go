package question

import (
	"context"
	"quizChallenge/models"
)

type Repository interface {
	GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error)
	Store(ctx context.Context, question *models.Question) error
	GetRandomQuestions(ctx context.Context, qType, category string, size int) ([]*models.Question, error)
}
