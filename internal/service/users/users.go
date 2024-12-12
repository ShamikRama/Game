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

func (r *UserService) GetUserStatus(userID int) (models.UserFullInfo, error) {
	userId, err := r.repo.GetUser(userID)
	if err != nil {
		return models.UserFullInfo{}, err
	}
	info := utils.StatusConverter(userId)
	return info, nil
}

func (r *UserService) GetUsersLeaders() ([]models.UserInfo, error) {
	leaders, err := r.repo.GetLeaders()
	if err != nil {
		return nil, err
	}
	info := utils.UserConverter(leaders)

	return info, nil
}
