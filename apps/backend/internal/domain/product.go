package domain

import (
	"context"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, p Product) (*Product, error)
	Update(ctx context.Context, p Product, id string) (*Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, p Product) (*Product, error)
	Update(ctx context.Context, p Product, id string) (*Product, error)
	Delete(ctx context.Context, id string) error
}
