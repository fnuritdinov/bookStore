package bookStore

import (
	"bookStore/cmd/internal/bookStore"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) AddBook(c fiber.Ctx) error {
	var b bookStore.Book

	if err := json.Unmarshal(c.Body(), &b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	bk, err := h.service.AddBook(b)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(bk)
}
