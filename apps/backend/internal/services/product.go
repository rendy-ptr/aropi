package service

import (
	"context"
	"errors"

	"github.com/rendy-ptr/aropi/backend/internal/domain"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *productService) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *productService) Create(ctx context.Context, p domain.Product) (*domain.Product, error) {
	if p.Price <= 0 {
		return nil, errors.New("harga harus lebih dari 0")
	}
	if p.Stock < 0 {
		return nil, errors.New("stok tidak boleh negatif")
	}
	return s.repo.Create(ctx, p)
}

func (s *productService) Update(ctx context.Context, p domain.Product, id string) (*domain.Product, error) {
	if p.Price <= 0 {
		return nil, errors.New("harga harus lebih dari 0")
	}
	if p.Stock < 0 {
		return nil, errors.New("stok tidak boleh negatif")
	}
	return s.repo.Update(ctx, p, id)
}

func (s *productService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
