package service

import (
	"crypto/sha1"
	"fmt"
	"web-programing-susu/pkg/models"
	"web-programing-susu/pkg/repository"
)

const salt = "hjgrhjqw124617ajfhajs"

// AuthService Сервис авторизациии
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	// Конструктор
	return &AuthService{repo: repository}
}
func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) generatePasswordHash(password string) string {
	// Метод хеширования пароля
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
