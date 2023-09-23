package service

import (
	"web-programing-susu/pkg/models"
	"web-programing-susu/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}
type Books interface {
	GetBooksPage(limit string, page string) (models.IndexViewData, error)
	CreateBook(book *models.Product) error
	EditBookPage(id string) (*models.Product, error)
	EditBookPost(id string, name string, author string, price string, rating string) error
	DeleteBook(id string) error
}
type Service struct {
	Authorization
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Books:         NewBookService(repos.Books),
	}
}
