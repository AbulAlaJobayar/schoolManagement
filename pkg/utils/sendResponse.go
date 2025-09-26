package utils

import "github.com/gofiber/fiber/v2"

// TApiResponse defines the common response format
type TApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// Meta holds pagination info
type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// SendResponse utility
func SendResponse(c *fiber.Ctx, data TApiResponse) error {
	responseData := TApiResponse{
		StatusCode: data.StatusCode,
		Success:    data.Success,
		Message:    data.Message,
		Meta:       data.Meta,
		Data:       data.Data,
	}

	return c.Status(data.StatusCode).JSON(responseData)
}
