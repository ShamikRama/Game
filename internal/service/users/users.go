package users

import (
	"Game/internal/models"
	"Game/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUsersService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) GetUserStatus(userID int) (models.User, error) {
	return r.repo.GetUser(userID)
}

func (r *UserService) GetUsersLeaders() ([]models.User, error) {
	return r.repo.GetLeaders()
}
