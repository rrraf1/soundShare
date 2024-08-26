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
	Link      string `json:"link"`
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

func (r *Repository) CreateMusic(context *fiber.Ctx) error {
	userID := context.Locals("userID").(uint)
	var music Music
	if err := context.BodyParser(&music); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Error creating music"})
	}

	music.UserID = int(userID)
	if err := r.DB.Create(&music).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Can't post music"})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Music post"})
	return nil
}

func (r *Repository) DeleteMusic(context *fiber.Ctx) error {
	var music Music
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID could not be empty"})
	}

	err := r.DB.Delete(&music, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not delete music"})
	}
	
	context.Status(http.StatusOK).JSON(&fiber.Map{"message" : "Book deleted"})
	return nil
}
