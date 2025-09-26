package auth

import "github.com/gofiber/fiber/v2"

func GetAllAuth(c *fiber.Ctx) error {
	return c.SendString("all auth")
	// return c.JSON(fiber.Map{"message": "All users"})
}