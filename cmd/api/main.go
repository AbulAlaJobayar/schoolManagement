package main

import (
	"log"
	"schoolmanagement/internal/router"
	"schoolmanagement/pkg/config"
	"schoolmanagement/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	// Connect DB
	db.ConnectDb()

	// Run migrations
	db.Migrate()
	app := fiber.New()

	api := app.Group("/api/v1")
	router.Routers(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!" + config.AppConfig.GoEnv)
	})
	log.Println("🚀 Server running on http://localhost:3000")
	log.Fatal(app.Listen(":" + config.AppConfig.Port))
}
