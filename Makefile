.PHONY: up down logs migrate run-producer run-consumer test build

# Start all infrastructure services
up:
	docker compose up -d

# Stop all infrastructure services
down:
	docker compose down

# Tail logs from all services
logs:
	docker compose logs -f

# Run database migrations
# TODO: Replace with actual migration tool (e.g., golang-migrate, goose)
migrate:
	@echo "TODO: Implement migration command"
	@echo "Example: migrate -path ./migrations -database postgres://streamsre:streamsre@localhost:5432/streamsre?sslmode=disable up"

# Run the producer locally
run-producer:
	go run ./cmd/producer

# Run the consumer locally
run-consumer:
	go run ./cmd/consumer

# Run all tests
test:
	go test ./...

# Build binaries
build:
	go build -o bin/producer ./cmd/producer
	go build -o bin/consumer ./cmd/consumer

# Tidy dependencies
tidy:
	go mod tidy

# Lint the codebase
lint:
	@echo "TODO: Add golangci-lint"

