package models

type Question struct {
	ID       interface{} `sql:"id" json:"id,omitempty" bson:"_id,omitempty"`
	Text     string      `sql:"text" json:"text" bson:"text"`
	Image    string      `sql:"image" json:"image" bson:"image"`
	Type     string      `sql:"type" json:"type" bson:"type"`
	Category string      `sql:"category" json:"category" bson:"category"`
	Answers  []Answer    `sql:"answers"  json:"answers" bson:"answers"`
}
