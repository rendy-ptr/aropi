package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	"github.com/rendy-ptr/aropi/backend/internal/middleware"
)

func registerProductRoutes(public fiber.Router, protected fiber.Router, c *container.Container, cfg *config.Config) {
	publicProduct := public.Group("/products")
	publicProduct.Get("/", c.Product.GetAll)

	protectedProduct := protected.Group("/products")
	protectedProduct.Post("/",
		middleware.RoleRequired("admin"),
		c.Product.Create,
	)
	protectedProduct.Put("/:id",
		middleware.RoleRequired("admin"),
		c.Product.Update,
	)
	protectedProduct.Delete("/:id",
		middleware.RoleRequired("admin"),
		c.Product.Delete,
	)
	protectedProduct.Get("/",
		middleware.RoleRequired("admin", "kasir"),
		c.Product.GetAll,
	)
}
