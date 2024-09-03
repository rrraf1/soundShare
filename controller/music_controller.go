package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rrraf1/soundshare/models"
	"gorm.io/gorm"
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
	var music models.Music
	if err := context.BodyParser(&music); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Error creating music"})
	}

	music.UserID = int(userID)
	if err := r.DB.Create(&music).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Can't post music"})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Music posted!"})
	return nil
}

func (r *Repository) DeleteMusic(context *fiber.Ctx) error {
	var music Music
	userID := context.Locals("userID").(uint)
	id := context.Params("id")

	if err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&music).Error; err != nil {
		return context.Status(http.StatusForbidden).JSON(&fiber.Map{"message": "Can't delete this music"})
	}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID could not be empty"})
	}

	err := r.DB.Delete(&music, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not delete music"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Music deleted"})
	return nil
}

func (r *Repository) UpdateMusic(context *fiber.Ctx) error {
	var input Music
	id := context.Params("id")

	if id == "" {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
	}

	if err := context.BodyParser(&input); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Error parsing request body"})
	}

	var existingMusic Music
	if err := r.DB.Where("id = ?", id).First(&existingMusic).Error; err != nil {
		return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Music not found"})
	}

	if err := r.DB.Model(&existingMusic).Updates(input).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not update music"})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Music updated"})
}

func (r *Repository) GetMusicByName(context *fiber.Ctx) error {
	var music Music

	if err := context.BodyParser(&music); err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Error creating music"})
	}

	MusicName := music.MusicName

	if err := r.DB.Where("music_name LIKE ?", "%"+MusicName+"%").First(&music).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "Music with that name not found"})
		}
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Error fetching music"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"data": music})
	return nil
}
