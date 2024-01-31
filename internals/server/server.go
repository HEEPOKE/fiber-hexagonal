package server

import (
	"log"

	"github.com/HEEPOKE/fiber-hexagonal/internals/app/handlers"
	"github.com/HEEPOKE/fiber-hexagonal/internals/app/services"
	"github.com/HEEPOKE/fiber-hexagonal/internals/core/interfaces"
	_ "github.com/HEEPOKE/fiber-hexagonal/pkg/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type Server struct {
	fib            *fiber.App
	accountHandler *handlers.AccountHandler
}

func NewServer(accountRepository interfaces.AccountRepositoryInterface) *Server {
	app := fiber.New(fiber.Config{
		ServerHeader:             "Fiber",
		AppName:                  "App v1.0",
		CaseSensitive:            true,
		StrictRouting:            true,
		Prefork:                  true,
		EnablePrintRoutes:        true,
		EnableSplittingOnParsers: true,
	})

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	accountService := services.NewAccountService(accountRepository)
	accountHandler := handlers.NewAccountHandler(*accountService)

	return &Server{
		fib:            app,
		accountHandler: accountHandler,
	}
}

func (s *Server) Init(address string) {
	s.routeConfig()

	err := s.fib.Listen(address)
	if err != nil {
		log.Fatalf("Failed To Start The Server: %v", err)
	}
}

func (s *Server) routeConfig() {
	apis := s.fib.Group("/apis")

	basicAuthMiddleware := basicauth.Config{
		Users: map[string]string{
			"admin": "        ",
		},
	}

	apis.Get("/docs/*", basicauth.New(basicAuthMiddleware), swagger.HandlerDefault)
	apis.Get("/monitor", monitor.New(monitor.Config{Title: "Monitor Page"}))

	account := apis.Group("/accounts")
	account.Get("/", s.accountHandler.GetListAccountAll)
}
