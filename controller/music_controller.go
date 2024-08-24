package controller

import (
	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
	"net/http"
)

type Music struct {
	ID        int    `json:"id"`
	MusicName string `json:"music_name"`
	Artist    string `json:"artist"`
	Genre     string `json:"genre"`
	UserID    int    `json:"user_id"`
}

func (r *Repository) GetMusics(context *fiber.Ctx) error {
	userID := context.Locals("userID").(uint)
	var music []Music
	if err := r.DB.Where("user_id = ?", userID).Find(&music).Error; err != nil {
        return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Error fetching music"})
    }
	if len(music) == 0 {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "User have not upload music yet"})
	}
	context.Status(http.StatusOK).JSON(music)
	return nil
}
