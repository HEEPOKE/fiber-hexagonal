package routes

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/handlers"
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/core/middleware"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutesAccount(app *fiber.App, db *gorm.DB) {
	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(*accountService)

	account := app.Group("/apis/accounts")
	account.Use(middleware.JwtMiddleware())

	account.Get("/", accountHandler.GetListAccountAll)
}
