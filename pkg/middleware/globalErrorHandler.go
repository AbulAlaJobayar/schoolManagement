package middleware

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	message := "Something went wrong"
	errorSources := []map[string]string{{"path": c.Path(), "message": message}}

	if ve, ok := err.(validator.ValidationErrors); ok {
		statusCode = fiber.StatusBadRequest
		message = "Validation failed"
		errorSources = []map[string]string{}
		for _, fe := range ve {
			errorSources = append(errorSources, map[string]string{
				"path":    fe.Field(),
				"message": fmt.Sprintf("%s is %s", fe.Field(), fe.Tag()),
			})
		}
	} else if appErr, ok := IsAppError(err); ok {
		statusCode = appErr.StatusCode
		message = appErr.Message
		errorSources = []map[string]string{
			{"path": c.Path(), "message": appErr.Err.Error()},
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		statusCode = fiber.StatusNotFound
		message = "Resource not found"
		errorSources = []map[string]string{{"path": c.Path(), "message": message}}
	} else if err != nil {
		message = err.Error()
		errorSources = []map[string]string{{"path": c.Path(), "message": message}}
	}

	log.Println(" Error stack:\n", string(debug.Stack()))

	return c.Status(statusCode).JSON(fiber.Map{
		"success":      false,
		"message":      message,
		"errorSources": errorSources,
		"err":          err.Error(),
	})
}
