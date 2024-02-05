package handlers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/helpers"
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// Get List All Accounts Data
// @Summary Get List All Accounts Data
// @Description Get List All Accounts Data
// @Tags Accounts
// @Accept application/json
// @Produce json
// @param Authorization header string true "Bearer token"
// @Router /accounts [get]
// @Success 200 {object} examples.SuccessAccountsGetAllResponse
// @Failure 400 {object} examples.FailedCommonResponse
func (ah *AccountHandler) GetListAccountAll(c *fiber.Ctx) error {
	accounts, err := ah.accountService.GetAccountsAll()
	if err != nil {
		return helpers.SendErrorResponse(c, err, constants.ACCOUNT_SERVICE, constants.ACCOUNT_LIST_ALL_FAILED)
	}

	statusMessage := response.StatusMessage{
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.ACCOUNT_SERVICE,
		Description: constants.ACCOUNT_LIST_ALL_SUCCESS,
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseMessage{
		Status:  statusMessage,
		Payload: accounts,
	})
}

// Get Accounts Profile
// @Summary Get Accounts Profile
// @Description Get Accounts Profile
// @Tags Accounts
// @Accept application/json
// @Produce json
// @param Authorization header string true "Bearer token"
// @Router /accounts/profile [get]
// @Success 200 {object} examples.SuccessAccountsProfileResponse
// @Failure 400 {object} examples.FailedCommonResponse
func (ah *AccountHandler) GetAccountProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accountID := claims["id"].(float64)

	accountData, err := ah.accountService.GetAccountDataWithId(accountID)
	if err != nil {
		return helpers.SendErrorResponse(c, err, constants.ACCOUNT_SERVICE, constants.GET_ACCOUNT_PROFILE_FAILED)
	}

	statusMessage := response.StatusMessage{
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.ACCOUNT_SERVICE,
		Description: constants.GET_ACCOUNT_PROFILE_SUCCESS,
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseMessage{
		Status:  statusMessage,
		Payload: accountData,
	})
}
