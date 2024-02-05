package repositories

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/response"
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

func (a *AccountRepository) GetAccountsAll() ([]*response.AccountResponseModel, error) {
	var accounts []*response.AccountResponseModel
	if err := a.conn.Select("id", "email", "username", "age", "is_active", "created_at", "updated_at").Preload("Blogs").Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *AccountRepository) GetAccountDataWithId(accountID float64) (*response.AccountResponseModel, error) {
	var accountData response.AccountResponseModel
	query := a.conn.Select("id", "email", "username", "age", "is_active", "created_at", "updated_at").
		Preload("Blogs").
		Where("id = ?", accountID).
		First(&accountData)

	if err := query.Error; err != nil {
		return nil, err
	}

	return &accountData, nil
}

func (a *AccountRepository) GetAccountDataWithEmail(email string) (*models.AccountModel, error) {
	var accountData models.AccountModel
	query := a.conn.Preload("Blogs").Where("email = ?", email).First(&accountData)

	if err := query.Error; err != nil {
		return nil, err
	}

	return &accountData, nil
}
