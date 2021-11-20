package api

import (
	"go-chat-service/src/controllers"
	"go-chat-service/src/dto"
	"go-chat-service/src/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddMessage(c *fiber.Ctx) error {
	body := new(dto.MessageDTO)
	body.ApplicationToken = c.Params("application_token")
	var parseErr error
	body.PerAppId, parseErr = strconv.Atoi(c.Params("chat_num"))
	if err := util.ParseBodyAndValidate(c, body); err != nil || parseErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			util.HttpResponse{
				StatusCode: http.StatusBadRequest,
				Data:       err,
			},
		)
	}
	res, err := controllers.AddMessage(body)
	if err != nil {
		return err
	}
	return c.Status(res.StatusCode).JSON(res.Data)
}
