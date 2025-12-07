package routes

import (
	"blogcms/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/posts", controllers.GetPosts)
	api.Post("/posts", controllers.CreatePost)
}
