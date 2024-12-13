package task

import (
	"database/sql"
	"fmt"
)

const (
	taskTable = "tasks"
	userTable = "users"
	refTable  = "referrals"
)

type TaskPsql struct {
	db *sql.DB
}

func NewTaskPsql(db *sql.DB) *TaskPsql {
	return &TaskPsql{
		db: db,
	}
}

func (r *TaskPsql) CompleteTask(userID int, goal_type string) error {
	const op = "sql.Task.Complete"

	query := fmt.Sprintf("INSERT INTO %s (user_id, goal) VALUES ($1, $2)", taskTable)
	_, err := r.db.Exec(query, userID, goal_type)
	if err != nil {
		return fmt.Errorf("%s: failed to complete task: %w", op, err)
	}

	return nil
}

func (r *TaskPsql) UpdatePoints(userID int, pointAdd int) error {
	const op = "sql.Task.UpdatePoints"

	query := fmt.Sprintf("UPDATE %s SET points = points + $1 WHERE id = $2", userTable)
	_, err := r.db.Exec(query, pointAdd, userID)
	if err != nil {
		return fmt.Errorf("%s: failed to update points: %w", op, err)
	}

	return nil
}

func (r *TaskPsql) CompleteRef(userID int, referrerID int) error {
	const op = "sql.Task.Complete"

	query := fmt.Sprintf("INSERT INTO %s (user_id, referrer_id) VALUES ($1, $2)", refTable)
	_, err := r.db.Exec(query, userID, referrerID)
	if err != nil {
		return fmt.Errorf("%s: failed to complete task: %w", op, err)
	}

	return nil
}

func (r *TaskPsql) UserExists(referrerID int) (bool, error) {
	const op = "sql.Auth.UserExists"

	query := "SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)"

	var exists bool
	err := r.db.QueryRow(query, referrerID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: failed to check user existence: %w", op, err)
	}

	return exists, nil
}
