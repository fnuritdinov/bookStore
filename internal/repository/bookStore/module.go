package bookStore

import (
	bookService "bookStore/internal/bookStore"
	"bookStore/pkg/db"
	"bookStore/pkg/logger"
)

type Repository interface {
	AddBook(book bookService.Book) (bookService.Book, error)
	GetBooks() ([]bookService.Book, error)
}

type repository struct {
	d db.DB
	l logger.Logger
}

func New(l logger.Logger, d db.DB) Repository {
	return &repository{
		l: l,
		d: d,
	}
}
