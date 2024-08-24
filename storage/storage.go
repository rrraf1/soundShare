package storage

import (
	"fmt"
	"log"

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

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
