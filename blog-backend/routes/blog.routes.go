package routes

import (
	"github.com/MishraShardendu22/controller"
	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(app *fiber.App) {
	blogGroup := app.Group("/api/blog")

	blogGroup.Post("/create", CreateBlogRoutes)
	blogGroup.Get("/getAll", GetAllBlogsRoutes)
	blogGroup.Get("/getByUser/:id", GetBlogsByUserRoutes)
	blogGroup.Get("/get/:id", GetBlogByIdRoutes)
	blogGroup.Delete("/delete/:id", DeleteBlogRoutes)
	blogGroup.Put("/update/:id", UpdateBlogRoutes)
}

func CreateBlogRoutes(c *fiber.Ctx) error {
	return controller.CreateBlog(c)
}

func GetAllBlogsRoutes(c *fiber.Ctx) error {
	return controller.GetAllBlogs(c)
}

func GetBlogsByUserRoutes(c *fiber.Ctx) error {
	return controller.GetBlogsByUser(c)
}

func GetBlogByIdRoutes(c *fiber.Ctx) error {
	return controller.GetBlogById(c)
}

func DeleteBlogRoutes(c *fiber.Ctx) error {
	return controller.DeleteBlog(c)
}

func UpdateBlogRoutes(c *fiber.Ctx) error {
	return controller.UpdateBlog(c)
}
