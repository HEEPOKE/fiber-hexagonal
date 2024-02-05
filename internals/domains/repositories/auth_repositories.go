package repositories

import (
	"errors"
	"time"

	"github.com/HEEPOKE/fiber-hexagonal/internals/app/utils"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/models/requests"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/constants"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthRepository struct {
	conn              *gorm.DB
	accountRepository *AccountRepository
}

func NewAuthRepository(conn *gorm.DB, accountRepository *AccountRepository) *AuthRepository {
	return &AuthRepository{
		conn:              conn,
		accountRepository: accountRepository,
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

func (a *AuthRepository) CreateTokens(request requests.LoginRequest) (accessToken string, refreshToken string, err error) {
	accountData, err := a.accountRepository.GetAccountDataWithEmail(request.Email)
	if err != nil {
		return "", "", err
	}

	checkPassword := utils.CheckPasswordHash(request.Password, *accountData.Password)
	if !checkPassword {
		return "", "", errors.New(constants.PASSWORD_NOT_CORRECT)
	}

	createToken := func(expirationHours time.Duration) (string, error) {
		claims := jwt.MapClaims{
			"id":       accountData.ID,
			"username": accountData.UserName,
			"email":    accountData.Email,
			"age":      accountData.Age,
			"exp":      time.Now().Add(expirationHours).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		return token.SignedString([]byte(configs.Cfg.SECRET_KEY))
	}

	accessToken, err = createToken(time.Hour * 72)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = createToken(time.Hour * 168)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
