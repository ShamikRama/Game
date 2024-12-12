package utils

import "Game/internal/models"

func LoginConvert(login models.Login) models.User {
	return models.User{
		Username: login.Username,
		Password: login.Password,
	}
}

func UserConverter(users []models.User) []models.UserInfo {
	info := make([]models.UserInfo, 0, len(users))
	for _, user := range users {
		info = append(info, models.UserInfo{
			Username: user.Username,
			Points:   user.Points,
		})
	}
	return info
}

func StatusConverter(user models.User) models.UserFullInfo {
	return models.UserFullInfo{
		Id:       user.Id,
		Username: user.Username,
		Points:   user.Points,
	}
}
