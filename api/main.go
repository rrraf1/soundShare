package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/adaptor/v2"
	"github.com/joho/godotenv"
	"github.com/rrraf1/soundshare/middleware"
	"github.com/rrraf1/soundshare/routes"
	"github.com/rrraf1/soundshare/storage"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
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
	
	repo := routes.NewRepository(db)
	repo.SetupRoutes(app)

	adaptor.FiberApp(app).ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(handler)))
}