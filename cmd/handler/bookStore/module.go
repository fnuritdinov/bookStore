package bookStore

import (
	"bookStore/internal/bookStore"
	"bookStore/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

type Handler interface {
	AddBook(c fiber.Ctx) error
	GetBooks(c fiber.Ctx) error
}

type handler struct {
	service bookStore.Service
	logger  logger.Logger
}

func New(l logger.Logger, service bookStore.Service) Handler {
	return &handler{
		service: service,
		logger:  l,
	}
}
