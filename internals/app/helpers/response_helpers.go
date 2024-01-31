package helpers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/gofiber/fiber/v2"
)

func HelperResponse(c *fiber.Ctx, StatusCode int, status response.StatusMessage, payload *interface{}) error {
	response := response.ResponseMessage{
		Status:  status,
		Payload: payload,
	}

	return c.Status(StatusCode).JSON(response)
}
