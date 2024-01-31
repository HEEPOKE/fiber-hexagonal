package services

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/core/interfaces"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
)

type AccountService struct {
	accountRepository interfaces.AccountRepositoryInterface
}

func NewAccountService(accountRepository interfaces.AccountRepositoryInterface) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (a *AccountService) GetAccountsAll() ([]*models.AccountModel, error) {
	return a.accountRepository.GetAccountsAll()
}
