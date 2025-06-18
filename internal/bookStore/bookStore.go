package bookStore

import (
	"bookStore/internal/domain"
	"context"
	"errors"
)

func (s *service) AddBook(book domain.Book) (domain.Book, error) {

	b, err := s.repo.AddBook(context.Background(), book)
	if err != nil {
		s.logger.Error("Can't add book errpr is - " + err.Error())
		return domain.Book{}, err
	}

	return b, nil
}

func (s *service) GetBooks() ([]domain.Book, error) {
	if len(s.books) == 0 {
		return nil, errors.New("no books found")
	}

	return s.books, nil
}
