package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	app = fiber.New()
	app.Get("/health", checkHealth)
}

func Serve() {
	log.Fatalln(app.Listen(":8000"))
}
