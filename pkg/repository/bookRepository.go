package repository

import (
	"database/sql"
	"strconv"
	"web-programing-susu/pkg/models"
)

type BooksRepository struct {
	db *sql.DB
}

func (b *BooksRepository) DeleteBook(id string) error {
	sqlText := "DELETE FROM" + " books WHERE id = " + id
	_, err := b.db.Exec(sqlText)
	return err
}

func (b *BooksRepository) EditBookPost(book *models.Product) error {
	sqlText := "UPDATE" + " books SET firstname = '" + book.Name + "', author = '" + book.Author + "', price = " + strconv.Itoa(book.Price) + ", rating = " + strconv.Itoa(book.Rating) + " WHERE id = " + strconv.Itoa(book.Id)
	_, err := b.db.Exec(sqlText)
	return err
}

func (b *BooksRepository) GetModelForEditPage(id string) (*models.Product, error) {
	sqlText := "SELECT * FROM" + " books " + "WHERE id = " + id
	row := b.db.QueryRow(sqlText)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Name, &prod.Author, &prod.Price, &prod.Rating)
	if err != nil {
		return nil, err
	}
	return &prod, nil
}

func (b *BooksRepository) CreateBook(book *models.Product) error {
	sqlText := "INSERT INTO" + " books (firstname, author, price, rating) VALUES ('" + book.Name + "', '" + book.Author + "', " + strconv.Itoa(book.Price) + "," + strconv.Itoa(book.Rating) + ")"
	_, err := b.db.Exec(sqlText)
	return err
}

func (b *BooksRepository) GetBooksPage(limit int, page int) ([]models.Product, error) {
	var sqlText = "SELECT * FROM" + " " + booksTable + " ORDER BY ID LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(page*limit)
	rows, err := b.db.Query(sqlText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Author, &p.Price, &p.Rating)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func NewBooksRepository(db *sql.DB) *BooksRepository {
	return &BooksRepository{db: db}
}
