package util

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseAPI(c *fiber.Ctx, status int, message string, data any, token ...string) error {
	response := fiber.Map{
		"data":    data,
		"message": message,
	}

	if len(token) > 0 {
		response["token"] = token[0]
	}

	c.Status(status).JSON(response)
	return nil
}
