package domain

import (
	"context"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]Category, error)
	FindByID(ctx context.Context, id string) (*Category, error)
	Create(ctx context.Context, c Category) (*Category, error)
	Update(ctx context.Context, id string, c Category) (*Category, error)
	Delete(ctx context.Context, id string) error
}

type CategoryService interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id string) (*Category, error)
	Create(ctx context.Context, c Category) (*Category, error)
	Update(ctx context.Context, id string, c Category) (*Category, error)
	Delete(ctx context.Context, id string) error
}
