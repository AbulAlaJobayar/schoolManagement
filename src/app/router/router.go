package router

import (
	"schoolmanagement/src/app/internal/auth"
	"schoolmanagement/src/app/internal/user"

	"github.com/gofiber/fiber/v2"
)

// Route defines a module route
type Route struct {
	Path  string
	Route func(router fiber.Router)
}

var modulesRoute = []Route{
	{
		Path:  "/user",
		Route: user.UserRoutes,
	},
	{
		Path:  "/auth",
		Route: auth.AuthRoutes,
	},
}

// SetupRoutes registers all routes to the app
func Routers(route fiber.Router) {
	// Register module routes
	for _, module := range modulesRoute {
		group := route.Group(module.Path)
		module.Route(group)
	}
}
