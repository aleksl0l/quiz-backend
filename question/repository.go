package question

import (
	"context"
	"quizChallenge/models"
)

type Repository interface {
	GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error)
}
