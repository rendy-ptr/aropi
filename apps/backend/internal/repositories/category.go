package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rendy-ptr/aropi/backend/internal/db"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
)

type categoryRepository struct {
	queries *db.Queries
}

func NewCategoryRepository(q *db.Queries) domain.CategoryRepository {
	return &categoryRepository{queries: q}
}

func (r *categoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	rows, err := r.queries.ListCategories(ctx)
	if err != nil {
		slog.Error("categoryRepository.FindAll", "error", err)
		return nil, err
	}

	categories := make([]domain.Category, len(rows))
	for i, row := range rows {
		categories[i] = domain.Category{
			ID:   row.ID.String(),
			Name: row.Name,
		}
	}

	return categories, nil
}

func (r *categoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		slog.Error("categoryRepository.FindByID: invalid uuid", "id", id, "error", err)
		return nil, err
	}

	row, err := r.queries.GetCategoryById(ctx, uuid)
	if err != nil {
		slog.Error("categoryRepository.FindByID", "id", id, "error", err)
		return nil, err
	}
	return &domain.Category{
		ID:   row.ID.String(),
		Name: row.Name,
	}, nil
}

func (r *categoryRepository) Create(ctx context.Context, c domain.Category) (*domain.Category, error) {
	row, err := r.queries.CreateCategory(ctx, c.Name)
	if err != nil {
		slog.Error("categoryRepository.Create", "name", c.Name, "error", err)
		return nil, err
	}
	return &domain.Category{
		ID:   row.ID.String(),
		Name: row.Name,
	}, nil
}

func (r *categoryRepository) Update(ctx context.Context, id string, c domain.Category) (*domain.Category, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		slog.Error("categoryRepository.Update: invalid uuid", "id", id, "error", err)
		return nil, err
	}

	row, err := r.queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:   uuid,
		Name: c.Name,
	})
	if err != nil {
		slog.Error("categoryRepository.Update", "id", id, "error", err)
		return nil, err
	}
	return &domain.Category{
		ID:   row.ID.String(),
		Name: row.Name,
	}, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		slog.Error("categoryRepository.Delete: invalid uuid", "id", id, "error", err)
		return err
	}

	_, err = r.queries.DeleteCategory(ctx, uuid)
	if err != nil {
		slog.Error("categoryRepository.Delete", "id", id, "error", err)
		return err
	}
	return nil
}
