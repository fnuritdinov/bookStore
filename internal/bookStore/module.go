package bookStore

import (
	"bookStore/internal/domain"
	"bookStore/internal/repository"
	"bookStore/pkg/db"
	"bookStore/pkg/logger"
	"context"
)

type Service interface {
	AddBook(book domain.Book) (domain.Book, error)
	GetBooks() ([]domain.Book, error)
}

type service struct {
	books  []domain.Book
	nextID int
	logger logger.Logger
	repo   repository.Repository
}

func New(ctx context.Context, l logger.Logger, dbURL string) Service {
	d := db.New(ctx, l, dbURL)
	r := repository.New(l, d)

	return &service{
		books:  make([]domain.Book, 0),
		nextID: 1,
		logger: l,
		repo:   r,
	}
}
