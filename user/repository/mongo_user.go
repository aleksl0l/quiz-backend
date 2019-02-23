package repository

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/models"
	"quizChallenge/user"
)

type mongoUserRepository struct {
	DB *mgo.Database
}

func NewMongoUserRepository(DB *mgo.Database) user.Repository {
	return &mongoUserRepository{DB}
}

func (r *mongoUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	u := &models.User{}
	err := r.DB.C("users").Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}
	return u, err
}

func (r *mongoUserRepository) Store(ctx context.Context, user *models.User) error {
	err := r.DB.C("users").Insert(user)
	return err
}
