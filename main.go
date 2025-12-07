package main

import (
	"fmt"
	"log"

	"blogcms/config"
	"blogcms/internal/models"
	"blogcms/internal/repository"
	"blogcms/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Load configuration
	config.InitConfig()

	//Connect to database
	repository.InitDB()
	repository.DB.AutoMigrate(&models.Post{})

	//Fiber initialization
	app := fiber.New()
	routes.SetUpRoutes(app)

	port := config.AppConfig.Server.Port
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
