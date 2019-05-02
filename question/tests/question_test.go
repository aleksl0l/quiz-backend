package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"quizChallenge/models"
	"quizChallenge/question/repository"
	"quizChallenge/tests"
	"testing"
)

func GetTestAnswerSample(isCorrect bool) models.Answer {
	return models.Answer{
		Text: tests.RandStringRunes(10),
		IsCorrect: isCorrect,
	}
}

func GetTestAnswersArraySample(n int) []models.Answer {
	answers := make([]models.Answer, 0, n)
	for i := 0; i < n - 1; i++ {
		answers = append(answers, GetTestAnswerSample(false))
	}
	answers = append(answers, GetTestAnswerSample(true))
	return answers
}

func GetTestQuestionSample(category string) *models.Question{
	return &models.Question{
		Text: "test_text",
		Type: "text",
		Category: category,
		Answers: GetTestAnswersArraySample(4),
	}
}

func TestQuestionStoreGetRepository(t *testing.T) {
	session, db := tests.GetTestDB()
	defer tests.FreeTestDB(session, db)

	asserts := assert.New(t)

	questionRepository := repository.NewMongoQuestionRepository(db)
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		question := GetTestQuestionSample("test_cat")
		err := questionRepository.Store(ctx, question)
		asserts.Equal(nil, err, "Store should return nil error")
	}

	qGet, err := questionRepository.GetQuestions(ctx, "text", "test_cat")
	asserts.Equal(nil, err, "Getting should return nil error")
	asserts.Equal(qGet[0].Category, "test_cat", "Username should equals")
	asserts.Equal(len(qGet), 10, "Length should be equal")
}

func TestQuestionRandomQuestionRepository(t *testing.T) {
	session, db := tests.GetTestDB()
	defer tests.FreeTestDB(session, db)

	asserts := assert.New(t)

	questionRepository := repository.NewMongoQuestionRepository(db)
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		question := GetTestQuestionSample("test_cat")

		err := questionRepository.Store(ctx, question)
		asserts.Equal(nil, err, "Store should return nil error")
	}

	qGet, err := questionRepository.GetRandomQuestions(ctx, "text", "test_cat", 5)
	asserts.Equal(nil, err, "Getting should return nil error")
	asserts.Equal(qGet[0].Category, "test_cat", "Username should equals")
	asserts.Equal(len(qGet), 5, "Length should be equal")
}