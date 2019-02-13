package models

type Answer struct {
	ID        rune   `sql:"id" json:"id"`
	Text      string `sql:"text" json:"text"`
	IsCorrect bool   `sql:"is_correct" json:"is_correct"`
}
