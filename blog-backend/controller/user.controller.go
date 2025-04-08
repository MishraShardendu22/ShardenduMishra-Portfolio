package controller

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/MishraShardendu22/database"
	"github.com/MishraShardendu22/model"
	"github.com/MishraShardendu22/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var Collection *mongo.Collection

func RegisterUser(c *fiber.Ctx) error {
	var newUser model.User

	if err := c.BodyParser(&newUser); err != nil {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Error parsing request body", nil)
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Image == "" {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Missing required fields: Name, Email, Image", nil)
	}

	if newUser.Password == "" {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Password is required", nil)
	}
	if len(newUser.Password) < 6 {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Password must be at least 6 characters long", nil)
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(newUser.Email) {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Invalid email format", nil)
	}

	coll := database.Client.Database("blog").Collection("users")

	filter := bson.M{"email": newUser.Email}
	existingUserResult := coll.FindOne(c.Context(), filter)

	if existingUserResult.Err() == nil {
		return util.ResponseAPI(c, fiber.StatusConflict, "User with this email already exists", nil)
	} else if existingUserResult.Err() != mongo.ErrNoDocuments {
		fmt.Println("Database error checking existing user:", existingUserResult.Err())
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Database error during user check", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Error processing registration", nil)
	}
	newUser.Password = string(hashedPassword)
	newUser.CreatedAt = time.Now().UTC()

	_, err = coll.InsertOne(c.Context(), newUser)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Error registering user", nil)
	}

	return util.ResponseAPI(c, fiber.StatusCreated, "User registered successfully", nil)
}

func GenerateJWT(email string, name string) (string, error) {
	var jwtSecret = os.Getenv("JWT_SECRET")
	fmt.Println("JWT_SECRET:", jwtSecret)
	if jwtSecret == "" {
		return "", fmt.Errorf("JWT_SECRET not set in environment")
	}

	claims := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func LoginUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Error parsing request body", nil)
	}

	if user.Email == "" || user.Password == "" {
		return util.ResponseAPI(c, fiber.StatusBadRequest, "Missing required fields: Email, Password", nil)
	}

	coll := database.Client.Database("blog").Collection("users")
	filter := bson.M{"email": user.Email}

	var existingUser model.User
	err := coll.FindOne(c.Context(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return util.ResponseAPI(c, fiber.StatusNotFound, "User not found", nil)
		}
		fmt.Println("Database error during login:", err)
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Database error", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return util.ResponseAPI(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	token, err := GenerateJWT(existingUser.Email, existingUser.Name)
	if err != nil {
		fmt.Println("Error generating JWT:", err)
		return util.ResponseAPI(c, fiber.StatusInternalServerError, "Error logging in", nil)
	}

	existingUser.Password = ""
	return util.ResponseAPI(c, fiber.StatusOK, "User logged in successfully", "", token)
}