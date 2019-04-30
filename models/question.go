package models

import "github.com/globalsign/mgo/bson"

type Question struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Text     string        `json:"text" bson:"text"`
	Image    string        `json:"image" bson:"image"`
	Type     string        `json:"type" bson:"type" validate:"required"`
	Category string        `json:"category" bson:"category" validate:"required"`
	Answers  []Answer      `json:"answers" bson:"answers" validate:"required,min=2"`
}
