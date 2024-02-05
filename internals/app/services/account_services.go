package services

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/core/interfaces"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
)

type AccountService struct {
	accountRepository interfaces.AccountRepositoryInterface
}

func NewAccountService(accountRepository interfaces.AccountRepositoryInterface) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (a *AccountService) GetAccountsAll() ([]*response.AccountResponseModel, error) {
	return a.accountRepository.GetAccountsAll()
}

func (a *AccountService) GetAccountDataWithId(accountID float64) (*response.AccountResponseModel, error) {
	return a.accountRepository.GetAccountDataWithId(accountID)
}

func (a *AccountService) GetAccountDataWithEmail(email string) (*models.AccountModel, error) {
	return a.accountRepository.GetAccountDataWithEmail(email)
}
