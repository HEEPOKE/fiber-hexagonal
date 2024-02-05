package interfaces

import "github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"

type AccountRepositoryInterface interface {
	GetAccountsAll() ([]*models.AccountModel, error)
	GetAccountDataWithId(accountID uint) (*models.AccountModel, error)
}
