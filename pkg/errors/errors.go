// pkg/errors/errors.go
package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AppError is the standardized error type
type AppError struct {
	Status  int         `json:"-"`              // HTTP status (not exposed in body directly)
	Code    string      `json:"code"`           // machine-friendly code (e.g. "user_not_found")
	Message string      `json:"message"`        // human-friendly message
	Details interface{} `json:"details,omitempty"` // optional (validation errors, field info)
	Err     error       `json:"-"`              // wrapped error (not exposed in response)
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Convenience constructors
func New(status int, code, msg string, details interface{}, err error) *AppError {
	return &AppError{Status: status, Code: code, Message: msg, Details: details, Err: err}
}
func BadRequest(code, msg string, details interface{}) *AppError {
	return New(fiber.StatusBadRequest, code, msg, details, nil)
}
func Unauthorized(code, msg string) *AppError {
	return New(fiber.StatusUnauthorized, code, msg, nil, nil)
}
func Forbidden(code, msg string) *AppError {
	return New(fiber.StatusForbidden, code, msg, nil, nil)
}
func NotFound(code, msg string) *AppError {
	return New(fiber.StatusNotFound, code, msg, nil, nil)
}
func Internal(msg string, err error) *AppError {
	return New(fiber.StatusInternalServerError, "internal_error", msg, nil, err)
}

// Wrap converts arbitrary error into AppError
func Wrap(err error) *AppError {
	if err == nil {
		return nil
	}
	var ae *AppError
	if errors.As(err, &ae) {
		return ae
	}

	// Fiber predefined errors
	if fe, ok := err.(*fiber.Error); ok {
		return New(fe.Code, strings.ToLower(strings.ReplaceAll(fe.Message, " ", "_")), fe.Message, nil, err)
	}

	// GORM errors
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return NotFound("record_not_found", "record not found")
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return Internal("invalid transaction", err)
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return BadRequest("missing_where_clause", "unsafe query without WHERE clause", nil)
	case errors.Is(err, gorm.ErrInvalidData):
		return BadRequest("invalid_data", "invalid data provided", nil)
	}

	// Default fallback
	return Internal("something went wrong", err)
}
