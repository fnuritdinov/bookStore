package bookStore

import "bookStore/pkg/logger"

type Service interface {
	AddBook(book Book) (Book, error)
	GetBooks() ([]Book, error)
}

type service struct {
	books  []Book
	nextID int
	logger logger.Logger
}

func New(l logger.Logger) Service {
	return &service{
		books:  make([]Book, 0),
		nextID: 1,
		logger: l,
	}
}
