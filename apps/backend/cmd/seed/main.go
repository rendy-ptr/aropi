package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	if err := seedAdmin(ctx, queries); err != nil {
		log.Fatal("failed to seed admin:", err)
	}

	log.Println("✅ Seeding completed!")
}

func seedAdmin(ctx context.Context, queries *db.Queries) error {
	email := "admin@aropi.com"

	existing, err := queries.GetUserByEmail(ctx, email)
	if err == nil && existing.Email == email {
		log.Println("⚠️  Admin already exists, skipping...")
		return nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:     "Admin",
		Email:    email,
		Password: string(hashed),
		Role:     db.UserRoleAdmin,
	})
	if err != nil {
		return err
	}

	log.Printf("✅ Admin created: %s (%s)\n", user.Name, user.Email)
	return nil
}
