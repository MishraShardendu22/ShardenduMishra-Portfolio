package main

import (
	"fmt"
	"os"

	"github.com/MishraShardendu22/database"
	"github.com/MishraShardendu22/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("This is the backend for my Blogging Application")

	app := fiber.New()

	err := godotenv.Load(".env")
	if err == nil {
		LoadEnvironmentVariables()
	} 
	
	SetUpCORS(app)
	ConnectDatabase(app)
	TestRoutes(app)
	SetUpRoutes(app)
	ListeningPort(app)
}

func SetUpRoutes(app *fiber.App) {
	// Define your application routes here.
}

func ListeningPort(app *fiber.App) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	fmt.Printf("Server is running on port %s\n", port)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

func TestRoutes(app *fiber.App) {
	app.Get("/test123", func(c *fiber.Ctx) error {
		return util.ResponseAPI(c, fiber.StatusOK, "Test123", nil)
	})
}

func ConnectDatabase(app *fiber.App) {
	MONGODB_URI := os.Getenv("MONGO_URI")
	if MONGODB_URI == "" {
		panic("MONGO_URI is not set in the environment variables")
	}
	database.ConnectToDatabase(MONGODB_URI)
}

func SetUpCORS(app *fiber.App) {
	fmt.Println("URL: ", os.Getenv("CLIENT_URL"))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CLIENT_URL"),
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
}

func LoadEnvironmentVariables() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file")
	}
}
