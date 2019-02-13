package question

import (
	"context"
	"quizChallenge/models"
)

type Usecase interface {
	GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error)
}
