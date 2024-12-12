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

// TODO cделать отдельную функцию конвертер
func (r *UserService) GetUsersLeaders() ([]models.UserInfo, error) {
	leaders, err := r.repo.GetLeaders()
	if err != nil {
		return nil, err
	}
	info := make([]models.UserInfo, 0, 5)
	for _, user := range leaders {
		info = append(info, models.UserInfo{
			Username: user.Username,
			Points:   user.Points,
		})
	}

	return info, nil
}

// TODO func (r *UserService) TaskTelegram()

// TODO func (r *UserService) TaskReferal()
