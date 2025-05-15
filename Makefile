include .env

migrate-up:
	@echo Running migration up...
	cd db/migrations && migrate -database $(DATABASE_DRIVER)://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSLMODE) -path . up

migrate-down:
	@echo Running migration down...
	cd db/migrations && migrate -database $(DATABASE_DRIVER)://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=$(DATABASE_SSLMODE) -path . down

run:
	@echo Running application...
	go run cmd/app/main.go

compose-up:
	@echo "Starting Docker containers..."
	docker compose up -d --build

compose-down:
	@echo "Stopping Docker containers..."
	docker compose down