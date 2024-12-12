package task

import (
	"database/sql"
	"fmt"
)

const (
	taskTable = "tasks"
	userTable = "users"
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

	query := fmt.Sprintf("UDATE %s SET points = points + $1 WHERE user_id = $2", userTable)
	_, err := r.db.Exec(query, pointAdd, userID)
	if err != nil {
		return fmt.Errorf("%s: failed to complete task: %w", op, err)
	}

	return nil
}
