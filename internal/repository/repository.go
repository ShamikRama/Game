package repository

import (
	"Game/internal/models"
	"Game/internal/repository/auth"
	"Game/internal/repository/users"
	"database/sql"
)

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPsql(db),
		User:          users.NewUsersPsql(db),
	}
}

type Authorization interface {
	Create(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type User interface {
	GetUser(UserID int) (models.User, error)
	GetLeaders() ([]models.User, error)
}
