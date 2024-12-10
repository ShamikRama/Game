package auth

import (
	"Game/internal/models"
	"database/sql"
)

type AuthPsql struct {
	db *sql.DB
}

func NewAuthPsql(db *sql.DB) *AuthPsql {
	return &AuthPsql{
		db: db,
	}
}

func (r *AuthPsql) Create(login models.Login) (int, error) {

}
