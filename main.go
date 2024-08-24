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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &storage.Config{
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		User:   os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
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
