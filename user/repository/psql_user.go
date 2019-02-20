package repository

import (
	"context"
	"database/sql"
	"quizChallenge/models"
	"quizChallenge/user"
)

type psqlUserRepository struct {
	DB *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPsqlUserRepository(DB *sql.DB) user.Repository {
	return &psqlUserRepository{DB}
}

func (r *psqlUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	row := r.DB.QueryRow(`SELECT id, username, password, email, real_name, city, age, game_points, is_disabled_ad
								FROM user_models WHERE username = $1`, username)
	u := new(models.User)
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Password,
		&u.Email,
		&u.RealName,
		&u.City,
		&u.Age,
		&u.GamePoints,
		&u.IsDisableAd,
	)
	return u, err
}

func (r *psqlUserRepository) Store(ctx context.Context, user *models.User) error {
	_, err := r.DB.Exec(
		`INSERT INTO user_models(username, email, password, real_name, city, age) VALUES($1, $2, $3, $4, $5, $6)`,
		user.Username,
		user.Email,
		user.Password,
		user.RealName,
		user.City,
		user.Age,
	)
	return err
}
