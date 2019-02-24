package models

import "time"

type Game struct {
	ID                interface{} `sql:"id" json:"id" bson:"_id,omitempty"`
	User1ID           interface{} `sql:"user1_id" json:"user1_id"`
	User2ID           interface{} `sql:"user2_id" json:"user2_id"`
	StartedAt         time.Time   `sql:"started_at" json:"started_at" bson:"startedAt,omitempty"`
	FinishedAt        *time.Time  `sql:"finished_at" json:"finished_at," bson:"finishedAt"`
	TypeQuestions     string      `sql:"type_questions" json:"type_questions" bson:"typeQuestions"`
	CategoryQuestions string      `sql:"category_questions" json:"category_questions" bson:"categoryQuestions,omitempty"`
}
