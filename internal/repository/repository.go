package repository

import (
	"Game/internal/models"
	"Game/internal/repository/auth"
	"database/sql"
)

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPsql(db),
	}
}

type Authorization interface {
	Create(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}
