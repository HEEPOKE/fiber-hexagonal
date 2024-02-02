package common

import (
	"fmt"
	"strings"

	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidatorCommon(c *fiber.Ctx, validator *validator.Validate, request interface{}, service, description string) error {
	if err := validator.Struct(request); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Validation failed on field '%s' - Tag: '%s', Value: '%v'", err.Field(), err.Tag(), err.Value())
			errorMessages = append(errorMessages, message)
		}

		concatenatedErrorMessage := strings.Join(errorMessages, "; ")

		responseData := response.StatusMessage{
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     service,
			Description: description,
			Error:       concatenatedErrorMessage,
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseMessage{
			Status:  responseData,
			Payload: nil,
		})
	}
	return nil
}
