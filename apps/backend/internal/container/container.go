package container

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rendy-ptr/aropi/backend/internal/config"
	"github.com/rendy-ptr/aropi/backend/internal/db"
	handler "github.com/rendy-ptr/aropi/backend/internal/handlers"
	repository "github.com/rendy-ptr/aropi/backend/internal/repositories"
	service "github.com/rendy-ptr/aropi/backend/internal/services"
)

type Container struct {
	User     *handler.UserHandler
	Product  *handler.ProductHandler
	Category *handler.CategoryHandler
	// Order     *handler.OrderHandler
	// Member    *handler.MemberHandler
	// Inventory *handler.InventoryHandler
	// Report    *handler.ReportHandler
}

func New(cfg *config.Config) *Container {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	queries := db.New(pool)

	userRepo := repository.NewUserRepository(queries)
	userSvc := service.NewUserService(userRepo, cfg.JWTSecret)
	userHdlr := handler.NewUserHandler(userSvc)

	categoryRepo := repository.NewCategoryRepository(queries)
	categorySvc := service.NewCategoryService(categoryRepo)
	categoryHdlr := handler.NewCategoryHandler(categorySvc)

	productRepo := repository.NewProductRepository(queries)
	productSvc := service.NewProductService(productRepo)
	productHdlr := handler.NewProductHandler(productSvc)

	// orderRepo := repository.NewOrderRepository(queries)
	// orderSvc := service.NewOrderService(orderRepo, productRepo)
	// orderHdlr := handler.NewOrderHandler(orderSvc)

	// memberRepo := repository.NewMemberRepository(queries)
	// memberSvc := service.NewMemberService(memberRepo)
	// memberHdlr := handler.NewMemberHandler(memberSvc)

	// inventoryRepo := repository.NewInventoryRepository(queries)
	// inventorySvc := service.NewInventoryService(inventoryRepo)
	// inventoryHdlr := handler.NewInventoryHandler(inventorySvc)

	// reportSvc := service.NewReportService(orderRepo, productRepo)
	// reportHdlr := handler.NewReportHandler(reportSvc)

	return &Container{
		User:     userHdlr,
		Product:  productHdlr,
		Category: categoryHdlr,
		// Order:     orderHdlr,
		// Member:    memberHdlr,
		// Inventory: inventoryHdlr,
		// Report:    reportHdlr,
	}
}
