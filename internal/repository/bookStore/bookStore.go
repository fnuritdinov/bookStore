package bookStore

import (
	bookService "bookStore/internal/bookStore"
	"errors"
)

type Book struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Price     float64  `json:"price"`
	Currency  string   `json:"currency"`
	Count     int      `json:"count"`
	Category  []string `json:"category"`
	Author    string   `json:"author"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	DeletedAt string   `json:"deleted_at"`
}

func (r *repository) AddBook(book bookService.Book) (bookService.Book, error) {
	return bookService.Book{}, errors.New("Not Implemented")
}

func (r *repository) GetBooks() ([]bookService.Book, error) {
	return nil, errors.New("Not Implemented")
}
