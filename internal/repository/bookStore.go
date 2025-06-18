package repository

import (
	"bookStore/internal/domain"
	"context"
	"errors"
)

func (r *repository) AddBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	query := "insert into books (name, author, category, count, price) values ($1, $2, $3, $4, $5)"
	_, err := r.d.Exec(ctx, query, book.Name, book.Author, book.Category, book.Count, book.Price)
	if err != nil {
		r.l.Error("Can't add book errpr is - " + err.Error())
		return domain.Book{}, err
	}

	return domain.Book{}, nil
}

func (r *repository) GetBooks() ([]domain.Book, error) {
	return nil, errors.New("Not Implemented")
}
