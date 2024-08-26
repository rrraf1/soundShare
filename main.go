package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rrraf1/soundshare/middleware"
	"github.com/rrraf1/soundshare/routes"
	"github.com/rrraf1/soundshare/storage"
)

func main() {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: error loading .env file:", err)
		}
	}

	db, err := storage.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all models
	if err := storage.MigrateAll(db); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(middleware.RecoveryMiddleware)
	r := routes.NewRepository(db)
	r.SetupRoutes(app)

	// Use PORT provided in environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}