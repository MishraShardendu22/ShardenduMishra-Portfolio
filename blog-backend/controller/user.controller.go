package controller

import (
	"fmt"
	"time"

	"github.com/MishraShardendu22/database"
	"github.com/MishraShardendu22/model"
	"github.com/MishraShardendu22/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func RegisterUser(c *fiber.Ctx) error {
	var newUser model.User

	if err := c.BodyParser(&newUser); err != nil {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Error parsing request body", nil)
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Image == "" || newUser.About == "" {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Missing required fields", nil)
	}

	coll := database.Client.Database("blog").Collection("users")

	filter := bson.M{"email": newUser.Email}
	existingUser := coll.FindOne(c.Context(), filter)

	if existingUser.Err() == nil {
		return util.ResponseAPI(c, fiber.StatusConflict, "User already exists", nil)
	} else if existingUser.Err() != mongo.ErrNoDocuments {
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Database error", nil)
	}

	newUser.CreatedAt = time.Now().UTC()

	if _, err := coll.InsertOne(c.Context(), newUser); err != nil {
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Error inserting user", nil)
	}

	return util.ResponseAPI(c, fiber.StatusOK, "User registered successfully", nil)
}

func LoginUser(c *fiber.Ctx) {
	fmt.Println("LoginUser")
}

func GetUserById(c *fiber.Ctx) {
	fmt.Println("GetUserById")
}

func UpdateUser(c *fiber.Ctx) {
	fmt.Println("UpdateUser")
}

func VerifyUser(c *fiber.Ctx) {
	fmt.Println("VerifyUser")
}
