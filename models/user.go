package models

import (
	"gorm.io/gorm"
)

type Users struct {
	ID       int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string   `gorm:"unique" json:"username"`
	Password string   `json:"password"`
	Musics   []Musics `gorm:"foreignKey:UserID" json:"musics"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
