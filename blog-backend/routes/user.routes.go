package routes

import (
	"github.com/MishraShardendu22/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/api/user")

	userGroup.Post("/login", LoginUserRoutes)
	userGroup.Get("/get/:id", GetUserByIdRoutes)
	userGroup.Put("/update/:id", UpdateUserRoutes)
	userGroup.Get("/verify/:id", VerifyUserRoutes)
	userGroup.Post("/register", RegisterUserRoutes)
}

func RegisterUserRoutes(c *fiber.Ctx) error {
	return controller.RegisterUser(c)
}

func LoginUserRoutes(c *fiber.Ctx) error {
	return controller.LoginUser(c)
}

func GetUserByIdRoutes(c *fiber.Ctx) error {
	return controller.GetUserById(c)
}

func UpdateUserRoutes(c *fiber.Ctx) error {
	return controller.UpdateUser(c)
}

func VerifyUserRoutes(c *fiber.Ctx) error {
	return controller.VerifyUser(c)
}