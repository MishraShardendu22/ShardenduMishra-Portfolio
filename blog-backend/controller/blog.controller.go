package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) {
	fmt.Println("CreateBlog")
}

func GetAllBlogs(c *fiber.Ctx) {
	fmt.Println("GetAllBlogs")

}

func GetBlogsByUser(c *fiber.Ctx) {
	fmt.Println("GetBlogsByUser")
}

func GetBlogById(c *fiber.Ctx) {
	fmt.Println("GetBlogById")
}

func DeleteBlog(c *fiber.Ctx) {
	fmt.Println("DeleteBlog")
}

func UpdateBlog(c *fiber.Ctx) {
	fmt.Println("UpdateBlog")
}
