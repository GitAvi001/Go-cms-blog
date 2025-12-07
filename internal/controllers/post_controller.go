package controllers

import (
	"blogcms/internal/models"
	"blogcms/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	repository.DB.Find(&posts)
	return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	repository.DB.Create(&post)
	return c.Status(fiber.StatusCreated).JSON(post)
}
