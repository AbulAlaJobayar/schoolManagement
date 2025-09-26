package user

import (
	"schoolmanagement/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	result := GetAllUsersServices("1")
	return utils.SendResponse(c, utils.TApiResponse{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "user retrieved Successfully",
		Data:       result,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result := GetUsersByIdServices(id)
	return utils.SendResponse(c, utils.TApiResponse{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "user retrieved Successfully",
		Data:       result,
	})
}
