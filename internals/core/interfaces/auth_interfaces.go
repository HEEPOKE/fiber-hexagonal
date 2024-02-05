package interfaces

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/requests"
)

type AuthRepositoryInterfaces interface {
	Register(newAccount *models.AccountModel) (*models.AccountModel, error)
	CreateTokens(request requests.LoginRequest) (accessToken string, refreshToken string, err error)
}
