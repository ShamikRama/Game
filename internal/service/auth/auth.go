package auth

import (
	"Game/internal/models"
	"Game/internal/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthorizationService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (r AuthService) Create(login models.Login) (int, error) {

	return id, nil
}
