package interfaces

import "github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"

type AuthRepositoryInterfaces interface {
	Register(newAccount *models.AccountModel) (*models.AccountModel, error)
}
