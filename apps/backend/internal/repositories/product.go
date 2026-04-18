package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rendy-ptr/aropi/backend/internal/db"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
)

type productRepository struct {
	queries *db.Queries
}

func NewProductRepository(q *db.Queries) domain.ProductRepository {
	return &productRepository{queries: q}
}

func (r *productRepository) FindAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	for _, row := range rows {
		products = append(products, domain.Product{
			ID:       row.ID.String(),
			Name:     row.Name,
			Price:    row.Price,
			Stock:    int(row.Stock),
			Category: row.Category.String,
		})
	}
	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		return nil, fmt.Errorf("invalid uuid: %w", err)
	}

	row, err := r.queries.GetProductById(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ID:       row.ID.String(),
		Name:     row.Name,
		Price:    row.Price,
		Stock:    int(row.Stock),
		Category: row.Category.String,
	}, nil
}

func (r *productRepository) Create(ctx context.Context, p domain.Product) (*domain.Product, error) {
	params := db.CreateProductParams{
		Name:  p.Name,
		Price: p.Price,
		Stock: int32(p.Stock),
		Category: pgtype.Text{
			String: p.Category,
			Valid:  p.Category != "",
		},
	}
	row, err := r.queries.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ID:       row.ID.String(),
		Name:     row.Name,
		Price:    row.Price,
		Stock:    int(row.Stock),
		Category: row.Category.String,
	}, nil
}

func (r *productRepository) Update(ctx context.Context, p domain.Product) (*domain.Product, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(p.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid uuid: %w", err)
	}
	params := db.UpdateProductParams{
		ID:    uuid,
		Name:  p.Name,
		Price: p.Price,
		Stock: int32(p.Stock),
		Category: pgtype.Text{
			String: p.Category,
			Valid:  p.Category != "",
		},
	}
	row, err := r.queries.UpdateProduct(ctx, params)
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ID:       row.ID.String(),
		Name:     row.Name,
		Price:    row.Price,
		Stock:    int(row.Stock),
		Category: row.Category.String,
	}, nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}
	_, err = r.queries.DeleteProduct(ctx, uuid)
	return err
}
