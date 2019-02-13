package models

type Question struct {
	ID      rune     `sql:"id" json:"id"`
	Text    string   `sql:"text" json:"text"`
	Image   string   `sql:"image" json:"image"`
	Answers []Answer `sql:"password"  json:"answers"`
}
