package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"schoolmanagement/internal/router"
	"schoolmanagement/pkg/config"
)

func main() {
	config.LoadConfig()
	app := fiber.New()

	api := app.Group("/api/v1")
	router.Routers(api)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!" + config.AppConfig.GoEnv)
	})

	log.Fatal(app.Listen(":" + config.AppConfig.Port))
}
