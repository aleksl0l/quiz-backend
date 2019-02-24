package models

type Answer struct {
	ID        interface{} `sql:"id" json:"id,omitempty" bson:"-"`
	Text      string      `sql:"text" json:"text" bson:"text"`
	IsCorrect bool        `sql:"is_correct" json:"is_correct" bson:"isCorrect"`
}
