package users

import (
	"Game/internal/models"
	"database/sql"
	"fmt"
)

const (
	userTable = "users"
)

type UserPsql struct {
	db *sql.DB
}

func NewUsersPsql(db *sql.DB) *UserPsql {
	return &UserPsql{
		db: db,
	}
}

func (r *UserPsql) GetUser(userID int) (models.User, error) {
	const op = "sql.Users.GetUser"

	var user models.User
	query := fmt.Sprintf("SELECT id, username, password_hash, points FROM %s WHERE id = $1", userTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&user.Id, &user.Username, &user.Password, &user.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("%s: user with id %d not found", op, userID)
		}
		return user, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return user, nil
}

func (r *UserPsql) GetLeaders() ([]models.User, error) {
	const op = "sql.Users.GetLeaders"

	users := make([]models.User, 0, 5)
	query := fmt.Sprintf("SELECT id, username, points FROM %s ORDER BY points DESC LIMIT 5", userTable)

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return users, fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query() // Исправлено: убрано передачу query в Query
	if err != nil {
		if err == sql.ErrNoRows {
			return users, fmt.Errorf("%s: no users found", op) // Исправлено: сообщение об ошибке
		}
		return users, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}
	defer rows.Close() // Добавлено: закрытие rows

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Points); err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	return users, nil
}
