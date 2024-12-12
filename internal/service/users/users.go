package users

import (
	"Game/internal/models"
	"Game/internal/repository"
	"Game/pkg/utils"
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

func (r *UserService) GetUsersLeaders() ([]models.UserInfo, error) {
	leaders, err := r.repo.GetLeaders()
	if err != nil {
		return nil, err
	}
	info := utils.UserConverter(leaders)

	return info, nil
}

// TODO func (r *UserService) TaskTelegram()

// TODO func (r *UserService) TaskReferal()
