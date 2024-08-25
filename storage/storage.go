package storage

import (
	// "fmt"
	"log"
	"os"

	"github.com/rrraf1/soundshare/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host, Port, User, DBName string
}

func MigrateAll(db *gorm.DB) error {
	if err := models.MigrateUsers(db); err != nil {
		log.Fatal(err)
		return err
	}
	if err := models.MigrateMusics(db); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func NewConnection() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}