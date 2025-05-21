package bookStore

import (
	"bookStore/internal/bookStore"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	AddBook(c fiber.Ctx) error
	GetBooks(c fiber.Ctx) error
}

type handler struct {
	service bookStore.Service
}

func New(service bookStore.Service) Handler {
	return &handler{service: service}
}
