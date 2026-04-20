package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/container"
	// "github.com/rendy-ptr/aropi/backend/internal/middleware"
)

func registerUserRoutes(public fiber.Router, protected fiber.Router, c *container.Container, cfg *config.Config) {
	publicUser := public.Group("/users")
	publicUser.Post("/login", c.User.Login)
	publicUser.Post("/register", c.User.Register)
	publicUser.Post("/logout", c.User.Logout)

	// protectedUser := protected.Group("/users")
	// protectedUser.Post("/",
	// 	middleware.RoleRequired("admin"),
	// 	c.User.Create,
	// )
	// protectedUser.Put("/:id",
	// 	middleware.RoleRequired("admin"),
	// 	c.User.Update,
	// )
	// protectedUser.Delete("/:id",
	// 	middleware.RoleRequired("admin"),
	// 	c.User.Delete,
	// )
	// protectedUser.Get("/",
	// 	middleware.RoleRequired("admin", "kasir"),
	// 	c.User.GetAll,
	// )
}
