package models

import (
	"gorm.io/gorm"
)

type Music struct {
    ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
    MusicName string `json:"music_name"`
    Artist    string `json:"artist"`
    Genre     string `json:"genre"`
    UserID    int    `json:"user_id"`
    Link      string `json:"link"`
}

func MigrateMusics(db *gorm.DB) error {
	err := db.AutoMigrate(&Music{})
	return err
}
