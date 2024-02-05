package handlers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/requests"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register
// @Summary register account
// @Description register account
// @Tags Auth
// @Accept application/json
// @Produce json
// @param Body body requests.RegisterRequest true "Register Request"
// @Router /auth/register [post]
// @Success 200 {object} examples.SuccessRegisterAccountResponse
// @Failure 400 {object} examples.FailedCommonResponse
func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var DataRequest requests.RegisterRequest
	if err := c.BodyParser(&DataRequest); err != nil {
		errorMessage := err.Error()
		responseData := response.StatusMessage{
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.AUTH_SERVICE,
			Description: constants.AUTH_REGISTER_ACCOUNT_FAILED,
			Error:       &errorMessage,
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseMessage{
			Status:  responseData,
			Payload: nil,
		})
	}

	isActive := true
	newAccount := models.AccountModel{
		Email:    DataRequest.Email,
		UserName: DataRequest.UserName,
		Password: &DataRequest.Password,
		Age:      DataRequest.Age,
		IsActive: &isActive,
	}

	result, err := ah.authService.Register(&newAccount)
	if err != nil {
		errorMessage := err.Error()
		responseData := response.StatusMessage{
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.AUTH_SERVICE,
			Description: constants.AUTH_REGISTER_ACCOUNT_FAILED,
			Error:       &errorMessage,
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseMessage{
			Status:  responseData,
			Payload: nil,
		})
	}

	statusMessage := response.StatusMessage{
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.AUTH_SERVICE,
		Description: constants.AUTH_REGISTER_ACCOUNT_SUCCESS,
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseMessage{
		Status:  statusMessage,
		Payload: result.CreatedAt,
	})
}
