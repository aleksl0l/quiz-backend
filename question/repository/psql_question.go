package repository

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"log"
	"quizChallenge/models"
	"quizChallenge/question"
)

type psqlQuestionRepo struct {
	DB *sql.DB
}

func NewPsqlQuestionRepository(db *sql.DB) question.Repository {
	return &psqlQuestionRepo{
		DB: db,
	}
}

func (r *psqlQuestionRepo) GetQuestions(ctx context.Context, qType, category string) ([]*models.Question, error) {
	prepareQuery, _ := r.DB.Prepare(
		`SELECT qm.id, qm.text, qm.image,
       		array_agg(am.id), array_agg(am.text), array_agg(am.is_correct)
			FROM question_models qm, answer_models am
			WHERE am.question = qm.id AND category=$1 AND type=$2
			GROUP BY qm.id, qm.text ORDER BY RANDOM() LIMIT 100`,
	)
	rows, _ := prepareQuery.Query(category, qType)
	questions := make([]*models.Question, 0, 100)
	tmp := &models.Question{Answers: make([]models.Answer, 4)}
	for rows.Next() {
		questionInstance := bindResponse(rows, tmp)
		questions = append(questions, questionInstance)
	}
	return questions, nil
}

func bindResponse(rows *sql.Rows, qm *models.Question) *models.Question {
	var text sql.NullString
	var image sql.NullString
	var idAnswers pq.Int64Array
	var textAnswers []string
	var isCorrectAnswers []bool
	err := rows.Scan(
		&qm.ID,
		&text,
		&image,
		&idAnswers,
		pq.Array(&textAnswers),
		pq.Array(&isCorrectAnswers),
	)
	if err != nil {
		log.Fatal(err)
	}
	qm.Text = text.String
	qm.Image = image.String
	for i, value := range idAnswers {
		qm.Answers[i].ID = rune(value)
		qm.Answers[i].Text = textAnswers[i]
		qm.Answers[i].IsCorrect = isCorrectAnswers[i]
	}
	return qm
}
