package routes

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/handlers"
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/repositories"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutesAccount(app *fiber.App, db *gorm.DB) {
	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(*accountService)

	account := app.Group("/apis/accounts")

	account.Use(jwtWare.New(jwtWare.Config{
		SigningKey: jwtWare.SigningKey{Key: []byte(configs.Cfg.SECRET_KEY)},
	}))
	account.Get("/", accountHandler.GetListAccountAll)
}
