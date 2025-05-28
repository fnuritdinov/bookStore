package main

import (
	bookHandler "bookStore/cmd/handler/bookStore"
	"bookStore/configs"
	bookService "bookStore/internal/bookStore"
	"bookStore/pkg/config"
	"bookStore/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func main() {

	app := fiber.New()

	l := logger.New("./app.log")

	service := bookService.New(l)
	handler := bookHandler.New(l, service)

	cfg := config.New("configs", "yml", "C:/Users/farrukh.nuritdinov/Desktop/bookStore/")
	srv := configs.NewService(cfg)

	baseApi := app.Group(fmt.Sprintf("/%s/%s", srv.Stage, srv.BaseURL))
	baseApi.Post("/books", handler.AddBook)
	baseApi.Get("/books", handler.GetBooks)

	l.Info("app started at " + srv.Stage + " port - " + strconv.Itoa(srv.Port))

	err := app.Listen(fmt.Sprintf(":%d", srv.Port))
	if err != nil {
		l.Error("app.Listen: Error - " + err.Error())
	}
}
