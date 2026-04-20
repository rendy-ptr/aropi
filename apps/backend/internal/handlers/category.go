package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/dto"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryHandler(s domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) GetAll(c fiber.Ctx) error {
	categories, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if categories == nil {
		return c.Status(200).JSON(dto.Response{
			Success: true,
			Message: "No categories found",
			Data:    []domain.Category{},
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
	})
}

func (h *CategoryHandler) GetByID(c fiber.Ctx) error {
	category, err := h.service.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Category retrieved successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) Create(c fiber.Ctx) error {
	bodyRequest := new(dto.CreateCategoryRequest)
	err := c.Bind().Body(bodyRequest)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
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
	category, err := h.service.Create(c.Context(), domain.Category{
		Name: bodyRequest.Name,
	})

	if err != nil {
		return c.Status(422).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(dto.Response{
		Success: true,
		Message: "Category created successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) Update(c fiber.Ctx) error {
	bodyRequest := new(dto.UpdateCategoryRequest)
	err := c.Bind().Body(bodyRequest)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
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
	category, err := h.service.Update(c.Context(), c.Params("id"), domain.Category{
		Name: bodyRequest.Name,
	})
	if err != nil {
		return c.Status(422).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Category updated successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) Delete(c fiber.Ctx) error {
	err := h.service.Delete(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Category deleted successfully",
	})
}
