package auth

import "github.com/gofiber/fiber/v2"

func AuthRoutes(app fiber.Router) {
app.Get("/", GetAllAuth)
}