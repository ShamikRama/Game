package auth

import (
	"Game/internal/models"
	"Game/internal/repository"
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 7 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (r *AuthService) Create(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return r.repo.Create(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
