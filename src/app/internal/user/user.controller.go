package user

import (
	"schoolmanagement/src/app/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	user := map[string]interface{}{
		"id":   1,
		"name": "Jobayar",
	}
	return  utils.SendResponse(c,utils.TApiResponse{
		StatusCode: fiber.StatusOK,
		Success: true,
		Message: "user retrieved Successfully",
		Data: user,
	})
	// return c.SendString("all users")
	// return c.JSON(fiber.Map{"message": "All users"})
}