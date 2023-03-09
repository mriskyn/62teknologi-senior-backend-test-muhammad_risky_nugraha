package response

import (
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/pkg/validation"
	"time"

	"github.com/gofiber/fiber/v2"
	// "gitlab.com/hmcorp/wallet-houston/internal/utils/translation"
	// "gitlab.com/hmcorp/wallet-houston/pkg/validation"
)

type Response struct {
	Service   string      `json:"service"`
	Message   string      `json:"message"`
	Status    bool        `json:"status"`
	ErrorCode string      `json:"error_code,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	CreatedAt int64       `json:"created_at,omitempty"`
}

type Data struct {
	Key   string
	Value interface{}
}

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	Count       int   `json:"count"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total,omitempty"`
	TotalPages  int64 `json:"total_pages,omitempty"`
}

func Success(c *fiber.Ctx, service string, message string, data ...*Data) error {
	var transformData interface{}

	if len(data) == 1 {
		transformData = data[0].Value
	} else if len(data) > 1 {
		populateData := make(map[string]interface{})
		for _, dt := range data {
			populateData[dt.Key] = dt.Value
		}
		transformData = populateData
	}

	return c.JSON(&Response{
		Service:   service,
		Message:   message,
		Status:    true,
		Data:      transformData,
		CreatedAt: time.Now().Unix(),
	})
}

func ErrResponse(c *fiber.Ctx, service string, StatusCode int, message string, err error) error {
	return c.Status(StatusCode).JSON(&Response{
		Service:   service,
		Message:   message,
		Status:    false,
		Error:     err.Error(),
		CreatedAt: time.Now().Unix(),
	})
}

func ErrValidationResponse(c *fiber.Ctx, service string, message string, validation []*validation.ErrorResponse) error {
	return c.Status(fiber.StatusBadRequest).JSON(&Response{
		Service:   service,
		Message:   message,
		Status:    false,
		Errors:    validation,
		CreatedAt: time.Now().Unix(),
	})
}
