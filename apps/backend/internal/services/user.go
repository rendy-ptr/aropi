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

func (s *userService) Login(ctx context.Context, u domain.User) (string, error) {
	user, err := s.repo.GetByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("userService.Login: user not found", "email", u.Email, "error", err)
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		slog.Error("userService.Login: password mismatch", "email", u.Email, "error", err)
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
		slog.Error("userService.Login: failed to sign token", "email", u.Email, "error", err)
		return "", errors.New("failed to generate token")
	}

	return tokenStr, nil
}

func (s *userService) Logout(ctx context.Context) error {
	return nil
}

func (s *userService) Register(ctx context.Context, u domain.User) (*domain.User, error) {
	if u.Name == "" {
		return nil, errors.New("name is required")
	}
	if u.Email == "" {
		return nil, errors.New("email is required")
	}
	if u.Password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("userService.Register: failed to hash password", "email", u.Email, "error", err)
		return nil, errors.New("failed to process registration")
	}

	user, err := s.repo.Register(ctx, domain.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
