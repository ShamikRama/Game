package service

import (
	"Game/internal/models"
	"Game/internal/repository"
	"Game/internal/service/auth"
)

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthorizationService(repo.Authorization),
	}
}

type Authorization interface {
	Create(user models.User) (int, error)
	GenerateJwtToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
