include .env

migrate-up:
	@echo Running migration up...
	cd db/migrations && migrate -database $(database.driver)://$(database.user):$(database.password)@$(database.host):$(database.port)/$(database.name)?sslmode=$(database.sslmode) -path . up

migrate-down:
	@echo Running migration down...
	cd db/migrations && migrate -database $(database.driver)://${database.user}:$(database.password)@$(database.host):$(database.port)/$(database.name)?sslmode=$(database.sslmode) -path . down

run:
	@echo Running application...
	go run cmd/app/main.go