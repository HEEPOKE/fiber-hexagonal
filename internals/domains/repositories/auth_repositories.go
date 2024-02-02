package repositories

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/utils"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	conn *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) *AuthRepository {
	return &AuthRepository{
		conn: conn,
	}
}

func (a *AuthRepository) Register(newAccount *models.AccountModel) (*models.AccountModel, error) {
	hashedPassword, err := utils.HashPassword(*newAccount.Password)
	if err != nil {
		return nil, err
	}
	newAccount.Password = &hashedPassword

	if err := a.conn.Create(newAccount).Error; err != nil {
		return nil, err
	}

	return newAccount, nil
}
