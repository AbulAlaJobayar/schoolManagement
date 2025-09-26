package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(router fiber.Router) {
	router.Get("/", GetAllUsers)
	router.Get("/:id",GetUserByID)
	// app.Post("/", CreateUser)
	// app.Put("/:id", UpdateUser)
	// app.Delete("/:id", DeleteUser)
}