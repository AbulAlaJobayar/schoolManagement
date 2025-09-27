package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "go.uber.org/zap"
	// "yourapp/pkg/logger"
)

type AppError struct {
	Status  int
	Code    string
	Message string
	Details map[string]string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

// ---------------------- Error Helpers ----------------------
func BadRequest(msg string, err error) *AppError {
	return &AppError{Status: http.StatusBadRequest, Code: "bad_request", Message: msg, Err: err}
}
func Internal(msg string, err error) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Code: "internal_error", Message: msg, Err: err}
}
func NotFound(msg string, err error) *AppError {
	return &AppError{Status: http.StatusNotFound, Code: "not_found", Message: msg, Err: err}
}
func ValidationError(err error) *AppError {
	if errs, ok := err.(validator.ValidationErrors); ok {
		details := make(map[string]string)
		for _, e := range errs {
			field := strings.ToLower(e.Field())
			details[field] = validationMessage(e)
		}
		return &AppError{
			Status:  http.StatusBadRequest,
			Code:    "validation_error",
			Message: "validation failed",
			Err:     err,
			Details: details,
		}
	}
	return BadRequest("invalid input", err)
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "must be at least " + e.Param() + " characters"
	case "max":
		return "must be at most " + e.Param() + " characters"
	default:
		return "is invalid"
	}
}

// ---------------------- Error Normalization ----------------------
func Wrap(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	// Fiber errors
	if fe, ok := err.(*fiber.Error); ok {
		return &AppError{Status: fe.Code, Code: "fiber_error", Message: fe.Message, Err: fe}
	}

	// GORM errors
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NotFound("record not found", err)
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return Internal("invalid transaction", err)
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return BadRequest("missing WHERE clause", err)
	case errors.Is(err, gorm.ErrInvalidData):
		return BadRequest("invalid data", err)
	}

	// Validator errors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return ValidationError(ve)
	}

	// Default fallback
	return Internal("something went wrong", err)
}

// ---------------------- Global Fiber ErrorHandler ----------------------
func ErrorHandler(c *fiber.Ctx, err error) error {
	appErr := Wrap(err)

	// // Logging
	// logger.Log.Error("request failed",
	// 	zap.String("path", c.Path()),
	// 	zap.String("method", c.Method()),
	// 	zap.Int("status", appErr.Status),
	// 	zap.String("code", appErr.Code),
	// 	zap.String("message", appErr.Message),
	// 	zap.Error(appErr.Err),
	// )

	// JSON Response
	return c.Status(appErr.Status).JSON(fiber.Map{
		"success":      false,
		"message":      appErr.Message,
		"errorSources": appErr.Details,
		"err":          appErr.Err.Error(),
		"timestamp":    time.Now().UTC(),
		"requestId":    c.Locals("requestid"),
	})
}
