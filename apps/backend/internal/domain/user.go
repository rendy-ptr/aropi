package domain

import (
	"context"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
	Register(ctx context.Context, u User) (*User, error)
}

type UserService interface {
	Login(ctx context.Context, u User) (token string, err error)
	Register(ctx context.Context, u User) (*User, error)
	Logout(ctx context.Context) error
}
