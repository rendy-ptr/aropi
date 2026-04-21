package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	"github.com/rendy-ptr/aropi/backend/internal/routes"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out interface{}) error {
	return v.validate.Struct(out)
}

func New(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:         "Aropi API v1",
		StructValidator: &structValidator{validate: validator.New()},
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Cookie"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
	}))

	c := container.New(cfg)

	routes.Register(app, c, cfg)

	return app
}
