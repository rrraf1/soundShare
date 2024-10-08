package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/rrraf1/soundshare/controller"
	"github.com/rrraf1/soundshare/middleware"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	rateLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit per client IP address
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests, please try again later.",
			})
		},
	})

	api := app.Group("/api", rateLimiter, middleware.AuthRequired)

	userRepo := controller.Repository{DB: r.DB}
	
	// Register controller
	app.Post("/register", userRepo.Register, rateLimiter)
	app.Post("/login", userRepo.Login, rateLimiter)
	
	// Music controller
	api.Get("/musics/all/:pages", userRepo.GetAllMusics)
	api.Get("/musics", userRepo.GetMusics)
	api.Post("/musics", userRepo.CreateMusic)
	api.Delete("/musics/:id", userRepo.DeleteMusic)
	api.Put("/musics/:id", userRepo.UpdateMusic)
	api.Get("/musics/details", userRepo.GetMusicByName)

	// // User controller
	api.Get("/users", userRepo.GetUsers)
}
