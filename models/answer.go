package models

import "github.com/globalsign/mgo/bson"

type Answer struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"-"`
	Text      string        `json:"text" bson:"text"`
	IsCorrect bool          `json:"is_correct" bson:"isCorrect"`
}
