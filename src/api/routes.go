package api

import (
	"log"

	"go-chat-service/src/config"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func Init() {
	app = fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", checkHealth)

	v1.Post("/applications/:application_token/chats", AddChat)
	v1.Post("/applications/:application_token/chats/:chat_num/messages", AddMessage)

}

func Serve() {
	log.Fatalln(app.Listen(":" + config.GetEnv("PORT", ":8000")))
}
