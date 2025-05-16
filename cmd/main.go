package main

import (
	bookHandler "bookStore/cmd/handler/bookStore"
	"bookStore/configs"
	bookService "bookStore/internal/bookStore"
	"bookStore/pkg/config"
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	service := bookService.New()
	handler := bookHandler.New(service)

	app.Post("/books", handler.AddBook)
	app.Get("/books", handler.GetBooks)

	cfg := config.New("configs", "yml", "C:/Users/farrukh.nuritdinov/Desktop/bookStore/")
	srv := configs.NewService(cfg)

	app.Listen(fmt.Sprintf(":%d", srv.Port))
	
}
