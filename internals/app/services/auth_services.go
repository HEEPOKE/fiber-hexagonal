package services

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/core/interfaces"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/requests"
)

type AuthService struct {
	authRepository interfaces.AuthRepositoryInterfaces
}

func NewAuthService(authRepository interfaces.AuthRepositoryInterfaces) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (a *AuthService) Register(newAccount *models.AccountModel) (*models.AccountModel, error) {
	return a.authRepository.Register(newAccount)
}

func (a *AuthService) CreateTokens(request requests.LoginRequest) (accessToken string, refreshToken string, err error) {
	return a.authRepository.CreateTokens(request)
}
