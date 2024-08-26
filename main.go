package main

import (
	"log"
	// "os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rrraf1/soundshare/middleware"
	"github.com/rrraf1/soundshare/routes"
	"github.com/rrraf1/soundshare/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
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

	app.Listen(":5000")
}
