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

	cfg := config.New("configs", "yml", "C:/Users/farrukh.nuritdinov/Desktop/bookStore/")
	srv := configs.NewService(cfg)

	baseApi := app.Group(fmt.Sprintf("/%s/%s", srv.Stage, srv.BaseURL))
	baseApi.Post("/books", handler.AddBook)
	baseApi.Get("/books", handler.GetBooks)

	app.Listen(fmt.Sprintf(":%d", srv.Port))

}
