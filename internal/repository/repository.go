package repository

import (
	"Game/internal/models"
	"Game/internal/repository/auth"
	"Game/internal/repository/task"
	"Game/internal/repository/users"
	"database/sql"
)

type Repository struct {
	Authorization
	User
	Task
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthPsql(db),
		User:          users.NewUsersPsql(db),
		Task:          task.NewTaskPsql(db),
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

type Task interface {
	CompleteTask(userID int, goal_type string) error
	UpdatePoints(userID int, pointAdd int) error
	CompleteRef(userID int, referrerID int) error
	UserExists(referrerID int) (bool, error)
}
