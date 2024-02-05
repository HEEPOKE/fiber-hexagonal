package interfaces

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
)

type AccountRepositoryInterface interface {
	GetAccountsAll() ([]*response.AccountResponseModel, error)
	GetAccountDataWithId(accountID float64) (*response.AccountResponseModel, error)
	GetAccountDataWithEmail(email string) (*models.AccountModel, error)
}
