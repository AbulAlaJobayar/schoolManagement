package auth

import "github.com/gofiber/fiber/v2"

func AuthRoutes(router fiber.Router) {
router.Get("/", GetAllAuth)
}