package service

import (
	"Game/internal/models"
	"Game/internal/repository"
	"Game/internal/service/auth"
	"Game/internal/service/users"
)

type Service struct {
	Authorization
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthorizationService(repo.Authorization),
		User:          users.NewUsersService(repo.User),
	}
}

type Authorization interface {
	Create(input models.Login) (int, error)
	GenerateJwtToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
	GetUserStatus(userId int) (models.User, error)
	GetUsersLeaders() ([]models.UserInfo, error)
}
