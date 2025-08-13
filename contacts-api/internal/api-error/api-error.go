package apiError

import "github.com/gofiber/fiber/v2"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewAPIError(code int, msg string) *APIError {
	return &APIError{
		Code:    code,
		Message: msg,
	}
}

func (e *APIError) Error() string {
	return e.Message
}

func AsAPIError(err error, target **APIError) bool {
	if e, ok := err.(*APIError); ok {
		*target = e
		return true
	}
	return false
}

func SendAPIError(c *fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(APIError{
		Code:    status,
		Message: message,
		Details: details,
	})
}