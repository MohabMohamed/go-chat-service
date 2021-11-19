package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func checkHealth(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"up": "ok"},
	)
}
