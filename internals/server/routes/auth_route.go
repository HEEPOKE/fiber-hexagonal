package routes

import (
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/handlers"
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/domains/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutesAuth(app *fiber.App, db *gorm.DB) {
	accountRepository := repositories.NewAccountRepository(db)

	authRepository := repositories.NewAuthRepository(db, accountRepository)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(*authService)

	auth := app.Group("/apis/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}
