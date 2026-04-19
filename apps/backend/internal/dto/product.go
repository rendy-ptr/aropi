package dto

type CreateProductRequest struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
}

type UpdateProductRequest struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
}
