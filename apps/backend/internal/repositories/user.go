package repository

import (
	"context"
	"log/slog"

	"github.com/rendy-ptr/aropi/backend/internal/db"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
)

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(q *db.Queries) domain.UserRepository {
	return &userRepository{queries: q}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		slog.Error("userRepository.GetByEmail", "email", email, "error", err)
		return nil, err
	}

	return &domain.User{
		ID:       row.ID.String(),
		Name:     row.Name,
		Email:    row.Email,
		Password: row.Password,
		Role:     string(row.Role),
	}, nil
}

func (r *userRepository) Register(ctx context.Context, name string, email string, password string) (*domain.User, error) {
	params := db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     db.UserRoleUser,
	}
	row, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		slog.Error("userRepository.Register", "email", email, "error", err)
		return nil, err
	}

	return &domain.User{
		ID:       row.ID.String(),
		Name:     row.Name,
		Email:    row.Email,
		Password: row.Password,
		Role:     string(row.Role),
	}, nil
}
