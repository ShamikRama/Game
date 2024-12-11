package auth

import (
	"Game/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type AuthPsql struct {
	db *sql.DB
}

func NewAuthPsql(db *sql.DB) *AuthPsql {
	return &AuthPsql{
		db: db,
	}
}

const (
	userTable = "users"
)

func (r *AuthPsql) Create(user models.User) (int, error) {
	const op = "sql.Auth.CreateUser"

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES($1, $2) RETURNING id", userTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var id int

	err = stmt.QueryRow(user.Username, user.Password).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return 0, fmt.Errorf("%s: username already exists: %w", op, err)
		}
		return 0, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return id, nil
}

func (r *AuthPsql) GetUser(username, password string) (models.User, error) {
	const op = "sql.Auth.GetUser"

	var user models.User
	query := fmt.Sprintf("SELECT id, username, password_hash, points FROM %s WHERE username = $1 AND password_hash = $2", userTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, password).Scan(&user.Id, &user.Username, &user.Password, &user.Points)
	if err != nil {
		return user, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return user, nil
}
