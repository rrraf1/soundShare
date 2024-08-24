package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rrraf1/soundshare/controller"
)


func AuthRequired(context *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {log.Fatal("Error loading .env file")}
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
    tokenString := context.Get("Authorization")
    if tokenString == "" {
        return context.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Token not found"})
    }
    
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")
    claims := &controller.Claims{}
    
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    
    if err != nil || !token.Valid {
        return context.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Invalid or expired token"})
    }
    
    // Store both username and userID in context
    context.Locals("username", claims.Username)
    context.Locals("userID", claims.UserID)
    
    return context.Next()
}