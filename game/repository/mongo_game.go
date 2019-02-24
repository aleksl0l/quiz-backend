package repository

import (
	"context"
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"quizChallenge/game"
	"quizChallenge/models"
	"time"
)

type mongoGameRepository struct {
	DB *mgo.Database
}

func NewMongoGameRepository(DB *mgo.Database) game.Repository {
	return &mongoGameRepository{DB}
}

func (r *mongoGameRepository) SearchGame(
	ctx context.Context, userId, gameType, gameCategory string) (*models.Game, error) {
	game := &models.Game{}
	err := r.DB.C("games").Find(bson.M{
		"typeQuestions":     gameType,
		"categoryQuestions": gameCategory,
		"user2id":           nil,
	}).One(&game)
	if err != nil {
		game.TypeQuestions = gameType
		game.CategoryQuestions = gameCategory
		game.User1ID = userId
		game.StartedAt = time.Now()
		err := r.DB.C("games").Insert(game)
		if err != nil {
			return nil, err
		}
	} else {
		gameUserId := game.User1ID.(string)
		if gameUserId == userId {
			return nil, errors.New("you already search game")
		}
		err = r.DB.C("games").Update(bson.M{"_id": game.ID}, bson.M{"$set": bson.M{"user2id": userId}})
		if err != nil {
			return nil, err
		}
		game.User2ID = userId
	}
	return game, nil
}

func (r *mongoGameRepository) GetGames(ctx context.Context, userId string) ([]*models.Game, error) {
	games := make([]*models.Game, 0, 10)
	err := r.DB.C("games").Find(bson.M{"$or": []bson.M{{"user1id": userId}, {"user2id": userId}}}).All(&games)
	return games, err
}
