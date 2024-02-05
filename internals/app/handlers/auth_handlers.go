package handlers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/helpers"
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
	var dataRequest requests.RegisterRequest
	if err := c.BodyParser(&dataRequest); err != nil {
		return helpers.SendErrorResponse(c, err, constants.AUTH_SERVICE, constants.AUTH_REGISTER_ACCOUNT_FAILED)
	}

	isActive := true
	newAccount := models.AccountModel{
		Email:    dataRequest.Email,
		UserName: dataRequest.UserName,
		Password: &dataRequest.Password,
		Age:      dataRequest.Age,
		IsActive: &isActive,
	}

	result, err := ah.authService.Register(&newAccount)
	if err != nil {
		return helpers.SendErrorResponse(c, err, constants.AUTH_SERVICE, constants.AUTH_REGISTER_ACCOUNT_FAILED)
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

// Login
// @Summary login
// @Description login
// @Tags Auth
// @Accept application/json
// @Produce json
// @param Body body requests.LoginRequest true "Login Request"
// @Router /auth/login [post]
// @Success 200 {object} examples.SuccessLoginResponse
// @Failure 400 {object} examples.FailedCommonResponse
func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var loginRequest requests.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return helpers.SendErrorResponse(c, err, constants.AUTH_SERVICE, constants.AUTH_LOGIN_ACCOUNT_FAILED)
	}

	accessToken, refreshToken, err := ah.authService.CreateTokens(loginRequest)
	if err != nil {
		return helpers.SendErrorResponse(c, err, constants.AUTH_SERVICE, constants.AUTH_LOGIN_ACCOUNT_FAILED)
	}

	statusMessage := response.StatusMessage{
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.AUTH_SERVICE,
		Description: constants.AUTH_LOGIN_ACCOUNT_SUCCESS,
	}

	result := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseMessage{
		Status:  statusMessage,
		Payload: result,
	})
}
