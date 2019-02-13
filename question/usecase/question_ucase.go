package usecase

import (
	"context"
	"quizChallenge/models"
	"quizChallenge/question"
	"time"
)

type questionUsecase struct {
	questionRepo   question.Repository
	contextTimeout time.Duration
}

func NewQuestionUsecase(q question.Repository, timeout time.Duration) question.Usecase {
	return &questionUsecase{
		questionRepo:   q,
		contextTimeout: timeout,
	}
}

func (u *questionUsecase) GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	res, err := u.questionRepo.GetQuestions(ctx, qType, category)
	if err != nil {
		return nil, err
	}
	return res, err

}
