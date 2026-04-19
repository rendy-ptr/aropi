package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/dto"
)

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(s domain.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if products == nil {
		return c.Status(200).JSON(dto.SuccessResponse{
			Success: true,
			Message: "No products found",
			Data:    []domain.Product{},
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	product, err := h.service.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var body dto.CreateProductRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Invalid body",
		})
	}
	product, err := h.service.Create(c.Context(), domain.Product{
		Name:     body.Name,
		Price:    body.Price,
		Stock:    body.Stock,
		Category: body.Category,
	})
	if err != nil {
		return c.Status(422).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {

	var body dto.UpdateProductRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Message: "Invalid body",
		})
	}
	product, err := h.service.Update(c.Context(), domain.Product{
		Name:     body.Name,
		Price:    body.Price,
		Stock:    body.Stock,
		Category: body.Category,
	}, c.Params("id"))
	if err != nil {
		return c.Status(422).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	err := h.service.Delete(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(204).JSON(dto.SuccessResponse{
		Success: true,
		Message: "Product deleted successfully",
	})
}
