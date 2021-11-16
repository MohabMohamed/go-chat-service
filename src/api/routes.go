package api

import (
	"log"

	"go-chat-service/src/config"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	app = fiber.New()
	app.Get("/health", checkHealth)
}

func Serve() {
	log.Fatalln(app.Listen(":" + config.GetEnv("PORT", ":8000")))
}
