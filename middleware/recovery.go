package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)


func RecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			
			c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"message": "Internal server error",
			})
		}
	}()
	return c.Next()
}
