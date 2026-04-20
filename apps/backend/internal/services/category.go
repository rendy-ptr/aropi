package service

import (
	"context"

	"github.com/rendy-ptr/aropi/backend/internal/domain"
)

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	return s.repo.FindAll(ctx)
}

func (s *categoryService) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *categoryService) Create(ctx context.Context, c domain.Category) (*domain.Category, error) {
	return s.repo.Create(ctx, c)
}

func (s *categoryService) Update(ctx context.Context, id string, c domain.Category) (*domain.Category, error) {
	return s.repo.Update(ctx, id, c)
}

func (s *categoryService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
