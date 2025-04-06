package routes

import (
	"github.com/MishraShardendu22/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/api/user")
	userGroup.Post("/register", RegisterUserRoutes)
}

func RegisterUserRoutes(c *fiber.Ctx) error {
	return controller.RegisterUser(c)
}
