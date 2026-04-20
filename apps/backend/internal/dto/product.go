package dto

type CreateProductRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Price    int64  `json:"price" validate:"required,min=0"`
	Stock    int    `json:"stock" validate:"required,min=0"`
	Category string `json:"category" validate:"required,min=3,max=100"`
}

type UpdateProductRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Price    int64  `json:"price" validate:"required,min=0"`
	Stock    int    `json:"stock" validate:"required,min=0"`
	Category string `json:"category" validate:"required,min=3,max=100"`
}
