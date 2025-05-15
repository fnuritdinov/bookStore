package main

import (
	bookHandler "bookStore/cmd/handler/bookStore"
	bookService "bookStore/internal/bookStore"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	service := bookService.New()
	handler := bookHandler.New(service)

	app.Post("/books", handler.AddBook)
	app.Get("/books", handler.GetBooks)

	app.Listen(":3000")
}
