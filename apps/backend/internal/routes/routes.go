package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	"github.com/rendy-ptr/aropi/backend/internal/middleware"
)

func Register(app *fiber.App, c *container.Container, cfg *config.Config) {
	api := app.Group("/api")

	public := api.Group("/public")

	protected := api.Group("/protected", middleware.AuthRequired(cfg.JWTSecret))

	// registerAuthRoutes(api, c, cfg)
	registerUserRoutes(public, protected, c, cfg)
	registerProductRoutes(public, protected, c, cfg)
	// registerOrderRoutes(api, c, cfg)
	// registerMemberRoutes(api, c, cfg)
	// registerInventoryRoutes(api, c, cfg)
	// registerReportRoutes(api, c, cfg)
}
