package service

import "Game/internal/models"

type Service struct {
	Authorization
}

func NewService() Service {
	return &Service{
		Authorization(),
	}
}

type Authorization interface {
	Create(user models.Login) (int, error)
}
