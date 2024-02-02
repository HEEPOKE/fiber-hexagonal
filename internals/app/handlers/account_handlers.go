package handlers

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/gofiber/fiber/v2"
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
		errorMessage := err.Error()
		responseData := response.StatusMessage{
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.ACCOUNT_SERVICE,
			Description: constants.ACCOUNT_LIST_ALL_FAILED,
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
		Service:     constants.ACCOUNT_SERVICE,
		Description: constants.ACCOUNT_LIST_ALL_SUCCESS,
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseMessage{
		Status:  statusMessage,
		Payload: accounts,
	})
}
