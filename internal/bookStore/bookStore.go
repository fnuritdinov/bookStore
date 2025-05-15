package bookStore

import (
	"errors"
	"time"
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

func (s *service) AddBook(book Book) (Book, error) {

	now := time.Now().Format("2006-01-02 15:04:05")
	book.ID = s.nextID
	b := Book{
		ID:        book.ID,
		Name:      book.Name,
		Price:     book.Price,
		Currency:  book.Currency,
		Category:  book.Category,
		Author:    book.Author,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.books = append(s.books, b)
	s.nextID++

	return b, nil
}

func (s *service) GetBooks() ([]Book, error) {
	if len(s.books) == 0 {
		return nil, errors.New("no books found")
	}

	return s.books, nil
}
