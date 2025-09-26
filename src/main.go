package main

import (
	"log"
	"schoolmanagement/src/app/pkg/config"
	"schoolmanagement/src/app/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	app := fiber.New()


	api := app.Group("/api/v1")
    router.Routers(api)
	
	// app.Use("/api", func(c *fiber.Ctx) error {
	// 	println("API route accessed:", c.OriginalURL())
	// 	return c.Next() // পরবর্তী route এ যাও
	// })

	app.Get("/api/user", func(c *fiber.Ctx) error {
		return c.SendString("User Route")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!" + config.AppConfig.GoEnv)
	})

	log.Fatal(app.Listen(":" + config.AppConfig.Port))
}
