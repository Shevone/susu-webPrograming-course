package repository

import (
	"database/sql"
	"web-programing-susu/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}
type Books interface {
	GetBooksPage(limit int, page int) ([]models.Product, error)
	CreateBook(book *models.Product) error
	GetModelForEditPage(id string) (*models.Product, error)
	EditBookPost(book *models.Product) error
	DeleteBook(id string) error
}
type Repository struct {
	Authorization
	Books
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Books:         NewBooksRepository(db),
	}
}
