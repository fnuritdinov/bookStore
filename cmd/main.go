package main

import (
	bookHandler "bookStore/cmd/handler/bookStore"
	"bookStore/configs"
	bookService "bookStore/internal/bookStore"
	"bookStore/pkg/config"
	"bookStore/pkg/logger"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func main() {

	app := fiber.New()

	l := logger.New("./app.log")
	cfg := config.New("configs", "yml", "D:/farukh/bookStore")

	var db configs.DB

	dbURL := db.URL(cfg)
	service := bookService.New(context.Background(), l, dbURL)
	handler := bookHandler.New(l, service)

	var srv = configs.Service{}
	srv.NewService(cfg)

	baseApi := app.Group(fmt.Sprintf("/%s/%s", srv.Stage, srv.BaseURL))
	baseApi.Post("/books", handler.AddBook)
	baseApi.Get("/books", handler.GetBooks)

	l.Info("app started at " + srv.Stage + " port - " + strconv.Itoa(srv.Port))

	err := app.Listen(fmt.Sprintf(":%d", srv.Port))
	if err != nil {
		l.Error("app.Listen: Error - " + err.Error())
	}
}
