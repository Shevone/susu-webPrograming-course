package service

import (
	"strconv"
	"web-programing-susu/pkg/models"
	"web-programing-susu/pkg/repository"
)

var (
	limitDb = 3
)

type BooksService struct {
	repo repository.Books
}

func (s *BooksService) DeleteBook(id string) error {
	return s.repo.DeleteBook(id)
}

func (s *BooksService) EditBookPost(id string, name string, author string, price string, rating string) error {
	idInt, _ := strconv.Atoi(id)
	priceInt, _ := strconv.Atoi(price)
	ratingInt, _ := strconv.Atoi(rating)
	return s.repo.EditBookPost(&models.Product{Id: idInt, Name: name, Author: author, Price: priceInt, Rating: ratingInt})
}

func (s *BooksService) EditBookPage(id string) (*models.Product, error) {
	return s.repo.GetModelForEditPage(id)
}

func (s *BooksService) CreateBook(book *models.Product) error {
	return s.repo.CreateBook(book)
}

func NewBookService(repository repository.Books) *BooksService {
	return &BooksService{repo: repository}
}
func (s *BooksService) GetBooksPage(limit string, page string) (models.IndexViewData, error) {
	// Вызов метода из репозитория для того чтобы получить список моделей
	l, p := GetPageAndLimit(limit, page)
	products, _ := s.repo.GetBooksPage(l, p)
	data := models.IndexViewData{BooksData: products, RealPage: strconv.Itoa(p), NextPage: strconv.Itoa(p + 1), PrevPage: strconv.Itoa(p - 1)}
	return data, nil
}

func GetPageAndLimit(limit string, page string) (int, int) {
	var resPage int
	var err error
	if limit != "" {
		if limitDb, err = strconv.Atoi(limit); err != nil {
			limitDb = 3
		}
	}
	if page != "" {
		if resPage, err = strconv.Atoi(page); err != nil {
			resPage = 0
		}
		if resPage < 0 {
			resPage = 0
		}
	} else {
		resPage = 0
	}

	return limitDb, resPage
}
