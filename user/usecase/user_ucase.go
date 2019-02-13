package usecase

import (
	"context"
	"quizChallenge/models"
	"quizChallenge/user"
	"time"
)

type userUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(u user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetByUsername(c context.Context, username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	res, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (u *userUsecase) Store(c context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	err := u.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
