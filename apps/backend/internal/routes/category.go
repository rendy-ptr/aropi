package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	"github.com/rendy-ptr/aropi/backend/internal/middleware"
)

func registerCategoryRoutes(public fiber.Router, protected fiber.Router, c *container.Container, cfg *config.Config) {
	publicCategory := public.Group("/categories")
	publicCategory.Get("/", c.Category.GetAll)

	protectedCategory := protected.Group("/categories")
	protectedCategory.Post("/",
		middleware.RoleRequired("admin"),
		c.Category.Create,
	)
	protectedCategory.Put("/:id",
		middleware.RoleRequired("admin"),
		c.Category.Update,
	)
	protectedCategory.Delete("/:id",
		middleware.RoleRequired("admin"),
		c.Category.Delete,
	)
	protectedCategory.Get("/",
		middleware.RoleRequired("admin", "kasir"),
		c.Category.GetAll,
	)

	protectedCategory.Get("/:id",
		middleware.RoleRequired("admin"),
		c.Category.GetByID,
	)
}
