APP_NAME = voucher-api
GO = go
DB_URL = $(shell cat .env | grep DATABASE_URL | cut -d '=' -f2-)

.PHONY: run
run:
	@echo "Starting the application..."
	@DATABASE_URL=$(DB_URL) $(GO) run ./cmd/main.go

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy

.PHONY: build
build:
	@echo "Building the application..."
	$(GO) build -o $(APP_NAME)

.PHONY: migrate
migrate:
	@echo "Running database migrations..."
	migrate -path ./migrations -database $(DB_URL) up

.PHONY: rollback
rollback:
	@echo "Rolling back database migrations..."
	migrate -path ./migrations -database $(DB_URL) down

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
