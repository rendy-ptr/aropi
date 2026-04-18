package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	"github.com/rendy-ptr/aropi/backend/internal/routes"
)

func New(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Aropi API v1",
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	c := container.New(cfg)

	routes.Register(app, c, cfg)

	return app
}
