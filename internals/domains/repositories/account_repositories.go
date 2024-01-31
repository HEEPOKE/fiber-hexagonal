package repositories

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	conn *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) *AccountRepository {
	return &AccountRepository{
		conn: conn,
	}
}

func (a *AccountRepository) GetAccountsAll() ([]*models.AccountModel, error) {
	var accounts []*models.AccountModel
	if err := a.conn.Select("id", "email", "username", "age", "is_active", "created_at", "updated_at").Preload("Blogs").Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}
