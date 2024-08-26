package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rrraf1/soundshare/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username, Password string
	ID                 uint
}

type Repository struct {
	DB *gorm.DB
}

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

func createToken(username string, userID uint) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	claims := &Claims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (r *Repository) Register(context *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			context.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}()

	var account User
	if err := context.BodyParser(&account); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Request failed"})
	}

	var existUser User
	if err := r.DB.Where("username = ?", account.Username).First(&existUser).Error; err == nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Username already exists"})
	} else if err != gorm.ErrRecordNotFound {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Database error"})
	}

	// Hash password
	hashedPass, err := HashPassword(account.Password)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Password hashing failed"})
	}
	account.Password = hashedPass

	// Create user
	if err := r.DB.Create(&account).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create user"})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "User created"})
}

func (r *Repository) Login(context *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			context.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}()

	var credentials User
	if err := context.BodyParser(&credentials); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Request failed"})
	}

	var user User
	if err := r.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Internal server error"})
	}

	if !VerifyPassword(user.Password, credentials.Password) {
		return context.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Password incorrect"})
	}

	// Use the new createToken function, passing both username and user ID
	tokenString, err := createToken(user.Username, user.ID)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Could not generate token"})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User found! Welcome!",
		"data": fiber.Map{
			"token": tokenString,
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
			},
		},
	})
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	var request struct {
		Username string `JSON:"username"`
	}

	if err := context.BodyParser(&request); err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Cannot parse body"})
	}

	var users User
	var musics []models.Music


	if err := r.DB.Where("username = ?", request.Username).First(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message" : "Cannot searh users"})
	}

	if err := r.DB.Where("user_id = ?", users.ID).Find(&musics).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message" : "Cannot searh users"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"Users": users, "Musics": musics})
	return nil
}
