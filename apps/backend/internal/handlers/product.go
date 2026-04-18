package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
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
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	if products == nil {
		return c.JSON([]domain.Product{})
	}
	return c.JSON(products)
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	product, err := h.service.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(product)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var body domain.Product
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid body"})
	}
	product, err := h.service.Create(c.Context(), body)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(201).JSON(product)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	var body domain.Product
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid body"})
	}
	product, err := h.service.Update(c.Context(), body)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(product)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	err := h.service.Delete(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(204).JSON(nil)
}
