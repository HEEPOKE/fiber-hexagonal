package server

import (
	"fmt"

	"github.com/HEEPOKE/fiber-hexagonal/internals/core/middleware"
	"github.com/HEEPOKE/fiber-hexagonal/internals/server/routes"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	_ "github.com/HEEPOKE/fiber-hexagonal/pkg/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

type Server struct {
	fib *fiber.App
	db  *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	version := fmt.Sprintf("App v%s", configs.Cfg.VERSION)
	app := fiber.New(fiber.Config{
		ServerHeader:             "Fiber",
		AppName:                  version,
		CaseSensitive:            true,
		StrictRouting:            true,
		Prefork:                  false,
		EnablePrintRoutes:        true,
		EnableSplittingOnParsers: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	return &Server{
		fib: app,
		db:  db,
	}
}

func (s *Server) Init() *fiber.App {
	basicAuthMiddleware := basicauth.Config{
		Users: map[string]string{
			"admin": "        ",
		},
	}

	apis := s.fib.Group("/apis")
	apis.Get("/docs/*", basicauth.New(basicAuthMiddleware), swagger.HandlerDefault)
	apis.Get("/monitor", basicauth.New(basicAuthMiddleware), monitor.New(monitor.Config{Title: "Monitor Page"}))

	apisAuth := s.fib.Group("/apis")
	apisAuth.Use(keyauth.New(middleware.ConfigDefault))
	apisAuth.Get("/mr", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	routes.SetupRoutesAuth(s.fib, s.db)
	routes.SetupRoutesAccount(s.fib, s.db)

	return s.fib
}
