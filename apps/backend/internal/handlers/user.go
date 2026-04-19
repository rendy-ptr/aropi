package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/dto"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var body dto.LoginRequest

	error := c.BodyParser(&body)
	if error != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Invalid body",
		})
	}

	token, err := h.service.Login(c.Context(), body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Login success",
		Data:    nil,
	})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body dto.CreateUserRequest
	error := c.BodyParser(&body)
	if error != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Invalid body" + error.Error(),
		})
	}

	user, err := h.service.Register(c.Context(), body.Name, body.Email, body.Password)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Register success",
		Data:    user,
	})
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	})
	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Logout success",
		Data:    nil,
	})
}
