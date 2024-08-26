package api

import (
	"log"
	"net/http"

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

// Handler - fungsi yang diekspor untuk Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := storage.NewConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := storage.MigrateAll(db); err != nil {
		log.Printf("Database migration error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app := fiber.New()
	app.Use(middleware.RecoveryMiddleware)
	
	repo := routes.NewRepository(db)
	repo.SetupRoutes(app)

	adaptor.FiberApp(app).ServeHTTP(w, r)
}