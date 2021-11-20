package api

import (
	"go-chat-service/src/controllers"
	"go-chat-service/src/dto"

	"github.com/gofiber/fiber/v2"
)

func AddChat(c *fiber.Ctx) error {
	body := new(dto.ChatDTO)
	body.ApplicationToken = c.Params("application_token")
	res, err := controllers.AddChat(body)
	if err != nil {
		return err
	}
	return c.Status(res.StatusCode).JSON(res.Data)
}
