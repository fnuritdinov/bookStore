package main

import (
	bookHandler "bookStore/cmd/handler/bookStore"
	bookService "bookStore/cmd/internal/bookStore"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	service := bookService.New()
	handler := bookHandler.New(service)

	app.Post("/books", handler.AddBook)

	app.Listen(":3000")
}
