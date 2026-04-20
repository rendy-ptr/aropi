package dto

type CreateProductRequest struct {
	Name       string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Price      int64  `json:"price" form:"price" validate:"required,min=0"`
	Stock      int    `json:"stock" form:"stock" validate:"required,min=0"`
	CategoryID string `json:"category_id" form:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	Name       string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Price      int64  `json:"price" form:"price" validate:"required,min=0"`
	Stock      int    `json:"stock" form:"stock" validate:"required,min=0"`
	CategoryID string `json:"category_id" form:"category_id" validate:"required,min=3,max=100"`
}
