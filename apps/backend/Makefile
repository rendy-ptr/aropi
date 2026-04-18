include .env
export

DB_URL=$(DATABASE_URL)
MIGRATION_PATH=internal/db/migration

run:
	go run ./cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -o aropi-api ./cmd/main.go

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down

migrate-drop:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" drop

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

sqlc:
	cd internal/db && sqlc generate

seed:
	go run ./cmd/seed/main.go

.PHONY: run build migrate-up migrate-down migrate-drop migrate-create sqlc seed
