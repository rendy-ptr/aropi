package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
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

	categoryIDs, err := seedCategories(ctx, queries)
	if err != nil {
		log.Fatal("failed to seed categories:", err)
	}

	if err := seedProducts(ctx, queries, categoryIDs); err != nil {
		log.Fatal("failed to seed products:", err)
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
		Name:     "Rendy",
		Email:    email,
		Password: string(hashed),
		Role:     db.UserRoleADMIN,
	})
	if err != nil {
		return err
	}

	log.Printf("✅ Admin created: %s (%s)\n", user.Name, user.Email)
	return nil
}

func seedCategories(ctx context.Context, queries *db.Queries) (map[string]pgtype.UUID, error) {
	categories := []string{
		"Coffee",
		"Non-Coffee",
		"Appetizer",
		"Dessert",
		"Food",
	}

	categoryIDs := make(map[string]pgtype.UUID)

	for _, name := range categories {
		existing, err := queries.GetCategoryByName(ctx, name)
		if err == nil {
			log.Printf("⚠️  Category '%s' already exists, skipping...\n", name)
			categoryIDs[name] = existing.ID
			continue
		}

		category, err := queries.CreateCategory(ctx, name)
		if err != nil {
			return nil, err
		}

		categoryIDs[name] = category.ID
		log.Printf("✅ Category created: %s\n", category.Name)
	}

	return categoryIDs, nil
}

func seedProducts(ctx context.Context, queries *db.Queries, categoryIDs map[string]pgtype.UUID) error {
	type productSeed struct {
		Name     string
		Price    int64
		Stock    int32
		Category string
		Image    string
	}

	products := []productSeed{
		// Coffee
		{Name: "Espresso", Price: 15000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?q=80&w=800"},
		{Name: "Americano", Price: 20000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1551030173-122aabc4489c?q=80&w=800"},
		{Name: "Cappuccino", Price: 25000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1534778101976-62847782c213?q=80&w=800"},
		{Name: "Cafe Latte", Price: 28000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1570968915860-54d5c301fa9f?q=80&w=800"},
		{Name: "Caramel Macchiato", Price: 30000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1485808191679-5f86510681a2?q=80&w=800"},

		// Non-Coffee
		{Name: "Matcha Latte", Price: 28000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1749280447307-31a68eb38673?q=80&w=800"},
		{Name: "Taro Milk", Price: 25000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1582570653002-8d512b39147a?q=80&w=800"},
		{Name: "Chocolate Ice", Price: 25000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1653122025865-5e75e63cf4ba?q=80&w=800"},
		{Name: "Lychee Tea", Price: 20000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1556679343-c7306c1976bc?q=80&w=800"},

		// Appetizer
		{Name: "French Fries", Price: 18000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1518013431117-eb1465fa5752?q=80&w=800"},
		{Name: "Onion Rings", Price: 20000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1639024471283-03518883512d?q=80&w=800"},
		{Name: "Nachos", Price: 25000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1776178393306-b42689a2548f?q=80&w=800"},
		{Name: "Chicken Wings", Price: 30000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1569691899455-88464f6d3ab1?q=80&w=800"},

		// Dessert
		{Name: "Pancake", Price: 25000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1528207776546-365bb710ee93?q=80&w=800"},
		{Name: "Waffle Ice Cream", Price: 30000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1648382721279-5e33912452e7?q=80&w=800"},
		{Name: "Cheesecake", Price: 35000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1533134242443-d4fd215305ad?q=80&w=800"},
		{Name: "Chocolate Brownie", Price: 28000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1606313564200-e75d5e30476c?q=80&w=800"},

		// Food
		{Name: "Nasi Goreng Spesial", Price: 35000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1603133872878-684f208fb84b?q=80&w=800"},
		{Name: "Spaghetti Bolognese", Price: 38000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1622973536968-3ead9e780960?q=80&w=800"},
		{Name: "Chicken Katsu Don", Price: 40000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1574484284002-952d92456975?q=80&w=800"},
		{Name: "Beef Burger", Price: 45000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?q=80&w=800"},
	}

	for _, p := range products {
		categoryID := categoryIDs[p.Category]
		existing, err := queries.GetProductByName(ctx, p.Name)
		if err == nil {
			// Update if exists to apply new image URLs
			_, err = queries.UpdateProduct(ctx, db.UpdateProductParams{
				ID:               existing.ID,
				ProductImageFile: p.Image,
				Name:             p.Name,
				Price:            p.Price,
				Stock:            p.Stock,
				CategoryID:       categoryID,
			})
			if err != nil {
				return err
			}
			log.Printf("✅ Product updated: %s\n", p.Name)
			continue
		}

		product, err := queries.CreateProduct(ctx, db.CreateProductParams{
			ProductImageFile: p.Image,
			Name:             p.Name,
			Price:            p.Price,
			Stock:            int32(p.Stock),
			CategoryID:       categoryID,
		})
		if err != nil {
			return err
		}

		log.Printf("✅ Product created: %s (Rp %d)\n", product.Name, product.Price)
	}

	return nil
}
