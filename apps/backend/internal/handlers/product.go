package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/dto"
)

type ProductHandler struct {
	service domain.ProductService
}

func NewProductHandler(s domain.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAll(c fiber.Ctx) error {
	products, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if products == nil {
		return c.Status(200).JSON(dto.Response{
			Success: true,
			Message: "No products found",
			Data:    []domain.Product{},
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    products,
	})
}

func (h *ProductHandler) GetByID(c fiber.Ctx) error {
	product, err := h.service.GetByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Product retrieved successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Create(c fiber.Ctx) error {
	bodyRequest := new(dto.CreateProductRequest)
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

	file, err := c.FormFile("product_image_file")
	if err != nil {
		return c.Status(400).JSON(dto.Response{
			Success: false,
			Message: "Image file is required",
			Error:   err.Error(),
		})
	}

	uniqueFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	destination := fmt.Sprintf("./public/uploads/products/%s", uniqueFileName)

	err = os.MkdirAll("./public/uploads/products", os.ModePerm)
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: "Failed to create directory structure",
			Error:   err.Error(),
		})
	}

	err = c.SaveFile(file, destination)
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: "Failed to save file",
			Error:   err.Error(),
		})
	}

	product, err := h.service.Create(c.Context(), domain.Product{
		ProductImageFile: uniqueFileName,
		Name:             bodyRequest.Name,
		Price:            bodyRequest.Price,
		Stock:            bodyRequest.Stock,
		CategoryID:       bodyRequest.CategoryID,
	})
	if err != nil {
		return c.Status(422).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(201).JSON(dto.Response{
		Success: true,
		Message: "Product created successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Update(c fiber.Ctx) error {
	bodyRequest := new(dto.UpdateProductRequest)
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

	file, err := c.FormFile("product_image_file")
	if err != nil {
		return c.Status(400).JSON(dto.Response{
			Success: false,
			Message: "Image file is required",
			Error:   err.Error(),
		})
	}

	uniqueFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	destination := fmt.Sprintf("./public/uploads/products/%s", uniqueFileName)

	err = os.MkdirAll("./public/uploads/products", os.ModePerm)
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: "Failed to create directory structure",
			Error:   err.Error(),
		})
	}

	err = c.SaveFile(file, destination)
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: "Failed to save file",
			Error:   err.Error(),
		})
	}

	product, err := h.service.Update(c.Context(), domain.Product{
		ProductImageFile: uniqueFileName,
		Name:             bodyRequest.Name,
		Price:            bodyRequest.Price,
		Stock:            bodyRequest.Stock,
		CategoryID:       bodyRequest.CategoryID,
	}, c.Params("id"))
	if err != nil {
		return c.Status(422).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Product updated successfully",
		Data:    product,
	})
}

func (h *ProductHandler) Delete(c fiber.Ctx) error {
	err := h.service.Delete(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(dto.Response{
			Success: false,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.Response{
		Success: true,
		Message: "Product deleted successfully",
	})
}
