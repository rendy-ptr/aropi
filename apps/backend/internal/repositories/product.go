package repository

import (
	"context"
	"log/slog"

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
		slog.Error("productRepository.FindAll", "error", err)
		return nil, err
	}
	var products []domain.Product
	for _, row := range rows {
		products = append(products, domain.Product{
			ID:               row.ID.String(),
			ProductImageFile: row.ProductImageFile,
			Name:             row.Name,
			Price:            row.Price,
			Stock:            int(row.Stock),
			CategoryID:       row.CategoryID.String(),
		})
	}
	return products, nil
}

func (r *productRepository) FindByID(ctx context.Context, id string) (*domain.Product, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		slog.Error("productRepository.FindByID: invalid uuid", "id", id, "error", err)
		return nil, err
	}

	row, err := r.queries.GetProductById(ctx, uuid)
	if err != nil {
		slog.Error("productRepository.FindByID", "id", id, "error", err)
		return nil, err
	}
	return &domain.Product{
		ID:               row.ID.String(),
		ProductImageFile: row.ProductImageFile,
		Name:             row.Name,
		Price:            row.Price,
		Stock:            int(row.Stock),
		CategoryID:       row.CategoryID.String(),
	}, nil
}

func (r *productRepository) Create(ctx context.Context, p domain.Product) (*domain.Product, error) {
	var uuid pgtype.UUID
	err := uuid.Scan(p.CategoryID)
	if err != nil {
		slog.Error("productRepository.Create: invalid uuid", "id", p.CategoryID, "error", err)
		return nil, err
	}
	params := db.CreateProductParams{
		ProductImageFile: p.ProductImageFile,
		Name:             p.Name,
		Price:            p.Price,
		Stock:            int32(p.Stock),
		CategoryID:       uuid,
	}
	row, err := r.queries.CreateProduct(ctx, params)
	if err != nil {
		slog.Error("productRepository.Create", "name", p.Name, "error", err)
		return nil, err
	}
	return &domain.Product{
		ID:               row.ID.String(),
		ProductImageFile: row.ProductImageFile,
		Name:             row.Name,
		Price:            row.Price,
		Stock:            int(row.Stock),
		CategoryID:       row.CategoryID.String(),
	}, nil
}

func (r *productRepository) Update(ctx context.Context, p domain.Product, id string) (*domain.Product, error) {
	var productUUID pgtype.UUID
	if err := productUUID.Scan(id); err != nil {
		slog.Error("productRepository.Update: invalid product uuid", "id", id, "error", err)
		return nil, err
	}

	var categoryUUID pgtype.UUID
	if err := categoryUUID.Scan(p.CategoryID); err != nil {
		slog.Error("productRepository.Update: invalid category uuid", "category_id", p.CategoryID, "error", err)
		return nil, err
	}

	params := db.UpdateProductParams{
		ID:               productUUID,
		ProductImageFile: p.ProductImageFile,
		Name:             p.Name,
		Price:            p.Price,
		Stock:            int32(p.Stock),
		CategoryID:       categoryUUID,
	}
	row, err := r.queries.UpdateProduct(ctx, params)
	if err != nil {
		slog.Error("productRepository.Update", "id", id, "error", err)
		return nil, err
	}
	return &domain.Product{
		ID:               row.ID.String(),
		ProductImageFile: row.ProductImageFile,
		Name:             row.Name,
		Price:            row.Price,
		Stock:            int(row.Stock),
		CategoryID:       row.CategoryID.String(),
	}, nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		slog.Error("productRepository.Delete: invalid uuid", "id", id, "error", err)
		return err
	}
	_, err = r.queries.DeleteProduct(ctx, uuid)
	if err != nil {
		slog.Error("productRepository.Delete", "id", id, "error", err)
		return err
	}
	return nil
}
