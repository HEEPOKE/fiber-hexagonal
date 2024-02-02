package common

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidatorCommon(c *fiber.Ctx, request interface{}, service, description string) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
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

	return nil
}
