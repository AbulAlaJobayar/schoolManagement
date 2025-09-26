package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(app fiber.Router) {
	app.Get("/", GetAllUsers)
	// app.Get("/:id", GetUserByID)
	// app.Post("/", CreateUser)
	// app.Put("/:id", UpdateUser)
	// app.Delete("/:id", DeleteUser)
}