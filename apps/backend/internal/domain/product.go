package domain

import (
	"context"
)



type Product struct {
	ID               string          `json:"id"`
	ProductImageFile string          `json:"product_image_file"`
	Name             string          `json:"name"`
	Price            int64           `json:"price"`
	Stock            int             `json:"stock"`
	Category         Category `json:"category"`
}

type ProductRepository interface {
	FindAll(ctx context.Context, search string, categoryID string) ([]Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, p Product) (*Product, error)
	Update(ctx context.Context, p Product, id string) (*Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductService interface {
	GetAll(ctx context.Context, search string, categoryID string) ([]Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, p Product) (*Product, error)
	Update(ctx context.Context, p Product, id string) (*Product, error)
	Delete(ctx context.Context, id string) error
}
