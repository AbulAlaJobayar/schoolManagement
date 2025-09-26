package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ValidateRequest middleware takes a struct to validate
func ValidateRequest(schema interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse body into the schema
		err := c.BodyParser(schema)
		if err != nil {
			return &AppError{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Invalid request body",
				Err:        err,
			}
		}

		// Run validation
		if err := validate.Struct(schema); err != nil {
			return err.(validator.ValidationErrors) // GlobalErrorHandler will catch
		}

		return c.Next()
	}
}
