package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/dto"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	bodyRequest := new(dto.LoginRequest)
	if err := c.Bind().Body(bodyRequest); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []fiber.Map
			for _, e := range validationErrors {
				errMsgs = append(errMsgs, fiber.Map{
					"field": e.Field(),
					"error": e.Error(),
				})
			}
			return c.Status(400).JSON(dto.Response{
				Success: false,
				Message: "Validation error",
				Error:   errMsgs,
			})
		}
		return c.Status(400).JSON(dto.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	token, err := h.service.Login(c.Context(), domain.User{
		Email:    bodyRequest.Email,
		Password: bodyRequest.Password,
	})
	if err != nil {
		return c.Status(401).JSON(dto.Response{
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

	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Login success",
		Data:    nil,
	})
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	bodyRequest := new(dto.CreateUserRequest)
	if err := c.Bind().Body(bodyRequest); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []fiber.Map
			for _, e := range validationErrors {
				errMsgs = append(errMsgs, fiber.Map{
					"field": e.Field(),
					"error": e.Error(),
				})
			}
			return c.Status(400).JSON(dto.Response{
				Success: false,
				Message: "Validation error",
				Error:   errMsgs,
			})
		}
		return c.Status(400).JSON(dto.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user, err := h.service.Register(c.Context(), domain.User{
		Name:     bodyRequest.Name,
		Email:    bodyRequest.Email,
		Password: bodyRequest.Password,
	})
	if err != nil {
		return c.Status(400).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Register success",
		Data:    user,
	})
}

func (h *UserHandler) Logout(c fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	})
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Logout success",
		Data:    nil,
	})
}
