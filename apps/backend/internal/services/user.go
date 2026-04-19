package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rendy-ptr/aropi/backend/internal/domain"
	"github.com/rendy-ptr/aropi/backend/internal/middleware"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo      domain.UserRepository
	jwtSecret string
}

func NewUserService(repo domain.UserRepository, jwtSecret string) domain.UserService {
	return &userService{repo: repo, jwtSecret: jwtSecret}
}

func (s *userService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("userService.Login: user not found", "email", email, "error", err)
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		slog.Error("userService.Login: password mismatch", "email", email, "error", err)
		return "", errors.New("invalid email or password")
	}

	claims := middleware.Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		slog.Error("userService.Login: failed to sign token", "email", email, "error", err)
		return "", errors.New("failed to generate token")
	}

	return tokenStr, nil
}

func (s *userService) Logout(ctx context.Context) error {
	return nil
}

func (s *userService) Register(ctx context.Context, name string, email string, password string) (*domain.User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("userService.Register: failed to hash password", "email", email, "error", err)
		return nil, errors.New("failed to process registration")
	}

	user, err := s.repo.Register(ctx, name, email, string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return user, nil
}
