package repository

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/models"
	"quizChallenge/question"
)

type mongoQuestionRepo struct {
	DB *mgo.Database
}

func NewMongoQuestionRepository(db *mgo.Database) question.Repository {
	return &mongoQuestionRepo{
		DB: db,
	}
}

func (r *mongoQuestionRepo) GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error) {
	questions := make([]*models.Question, 0, 20)
	err := r.DB.C("questions").Find(bson.M{"type": qType, "category": category}).Limit(20).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (u *mongoQuestionRepo) Store(ctx context.Context, question *models.Question) error {
	question.ID = bson.NewObjectId()
	err := u.DB.C("questions").Insert(question)
	if err != nil {
		return err
	}
	return nil
}
