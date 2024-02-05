package helpers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

func HelperResponse(c *fiber.Ctx, StatusCode int, status response.StatusMessage, payload *interface{}) error {
	response := response.ResponseMessage{
		Status:  status,
		Payload: payload,
	}

	return c.Status(StatusCode).JSON(response)
}

func SendErrorResponse(c *fiber.Ctx, err error, service, description string) error {
	errorMessage := err.Error()
	responseData := response.StatusMessage{
		Code:        constants.CODE_FAILED,
		Message:     constants.MESSAGE_FAIL,
		Service:     service,
		Description: description,
		Error:       &errorMessage,
	}

	return c.Status(fiber.StatusBadRequest).JSON(response.ResponseMessage{
		Status:  responseData,
		Payload: nil,
	})
}
