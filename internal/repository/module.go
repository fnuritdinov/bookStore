package repository

import (
	"bookStore/internal/domain"
	"bookStore/pkg/db"
	"bookStore/pkg/logger"
	"context"
)

type Repository interface {
	AddBook(ctx context.Context, book domain.Book) (domain.Book, error)
	GetBooks() ([]domain.Book, error)
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
