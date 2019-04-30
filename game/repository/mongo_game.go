package repository

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/game"
	"quizChallenge/models"
)

type mongoGameRepository struct {
	DB *mgo.Database
}

func NewMongoGameRepository(DB *mgo.Database) game.Repository {
	return &mongoGameRepository{DB}
}

func (r *mongoGameRepository) GetGamesByCondition(ctx context.Context, condition bson.M) (*models.Game, error) {
	game := &models.Game{}
	err := r.DB.C("games").Find(condition).One(&game)
	if err != nil {
		return nil, err
	}
	return game, err
}

func (r *mongoGameRepository) GetGames(ctx context.Context, userId string) ([]*models.Game, error) {
	games := make([]*models.Game, 0, 10)
	userObjectId := bson.ObjectIdHex(userId)
	err := r.DB.C("games").Find(bson.M{
		"$or": []bson.M{
			{"user1id": userObjectId},
			{"user2id": userObjectId},
		},
	}).All(&games)
	return games, err
}

func (r *mongoGameRepository) Store(ctx context.Context, game *models.Game) error {
	err := r.DB.C("games").Insert(game)
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoGameRepository) Update(ctx context.Context, game *models.Game, update bson.M) error {
	err := r.DB.C("games").Update(bson.M{"_id": game.ID}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (r *mongoGameRepository) GetGameByGameIdQuestionId(ctx context.Context, gameId, questionId string) (*models.Game, error) {
	gameObjectId := bson.ObjectIdHex(gameId)
	questionObjectId := bson.ObjectIdHex(questionId)
	game := &models.Game{}
	err := r.DB.C("games").Find(bson.M{"_id": gameObjectId, "questions._id": questionObjectId}).One(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}