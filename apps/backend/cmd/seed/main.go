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
		{Name: "Espresso", Price: 15000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=500&q=80"},
		{Name: "Americano", Price: 20000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1551030173-122aabc4489c?w=500&q=80"},
		{Name: "Cappuccino", Price: 25000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1534778101976-62847782c213?w=500&q=80"},
		{Name: "Cafe Latte", Price: 28000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1570968915860-54d5c301fa9f?w=500&q=80"},
		{Name: "Caramel Macchiato", Price: 30000, Stock: 100, Category: "Coffee", Image: "https://images.unsplash.com/photo-1485808191679-5f86510681a2?w=500&q=80"},

		// Non-Coffee
		{Name: "Matcha Latte", Price: 28000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1515823662972-da6a29051671?w=500&q=80"},
		{Name: "Taro Milk", Price: 25000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1626359055913-75c1328990c7?w=500&q=80"},
		{Name: "Chocolate Ice", Price: 25000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1553787762-b5f52896ec9c?w=500&q=80"},
		{Name: "Lychee Tea", Price: 20000, Stock: 100, Category: "Non-Coffee", Image: "https://images.unsplash.com/photo-1556679343-c7306c1976bc?w=500&q=80"},

		// Appetizer
		{Name: "French Fries", Price: 18000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1576107232684-1279f3908594?w=500&q=80"},
		{Name: "Onion Rings", Price: 20000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1639024471283-03518883512d?w=500&q=80"},
		{Name: "Nachos", Price: 25000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1513456811591-a1ed42d1cb15?w=500&q=80"},
		{Name: "Chicken Wings", Price: 30000, Stock: 50, Category: "Appetizer", Image: "https://images.unsplash.com/photo-1569691899455-88464f6d3ab1?w=500&q=80"},

		// Dessert
		{Name: "Pancake", Price: 25000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1528207776546-384cb1119671?w=500&q=80"},
		{Name: "Waffle Ice Cream", Price: 30000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1562376552-0d160a2f5dd3?w=500&q=80"},
		{Name: "Cheesecake", Price: 35000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1533134242443-d4fd215305ad?w=500&q=80"},
		{Name: "Chocolate Brownie", Price: 28000, Stock: 50, Category: "Dessert", Image: "https://images.unsplash.com/photo-1606313564200-e75d5e30476c?w=500&q=80"},

		// Food
		{Name: "Nasi Goreng Spesial", Price: 35000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1603133872878-684f208fb84b?w=500&q=80"},
		{Name: "Spaghetti Bolognese", Price: 38000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1622973536968-3ead9e780960?w=500&q=80"},
		{Name: "Chicken Katsu Don", Price: 40000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1574484284002-952d92456975?w=500&q=80"},
		{Name: "Beef Burger", Price: 45000, Stock: 50, Category: "Food", Image: "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=500&q=80"},
	}

	for _, p := range products {
		existing, err := queries.GetProductByName(ctx, p.Name)
		if err == nil {
			log.Printf("⚠️  Product '%s' already exists, skipping...\n", existing.Name)
			continue
		}

		categoryID := categoryIDs[p.Category]

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
