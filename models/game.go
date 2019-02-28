package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Game struct {
	ID                bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	User1ID           bson.ObjectId  `json:"user1_id"`
	User2ID           bson.ObjectId `json:"user2_id" bson:",omitempty"`
	StartedAt         *time.Time     `json:"started_at" bson:"startedAt,omitempty"`
	FinishedAt        *time.Time     `json:"finished_at," bson:"finishedAt"`
	TypeQuestions     string         `json:"type_questions" bson:"typeQuestions"`
	CategoryQuestions string         `json:"category_questions" bson:"categoryQuestions,omitempty"`
}
