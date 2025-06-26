# Makefile for AI Household Budget Microservice

.PHONY: help build run test clean docker-build docker-up docker-down docker-logs air install-deps

# Default target
help:
	@echo "Available commands:"
	@echo "  make help         - Show this help message"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application locally"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make install-deps - Install dependencies"
	@echo "  make air          - Run with hot reload (Air)"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make docker-logs  - Show Docker logs"

# Build the application
build:
	@echo "Building application..."
	go build -o bin/main ./cmd/main.go

# Run the application locally
run:
	@echo "Running application..."
	go run ./cmd/web_api/main.go

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf tmp/
	go clean

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Run with hot reload using Air
air:
	@echo "Starting Air for hot reload..."
	air

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker-compose build

docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down

docker-logs:
	@echo "Showing Docker logs..."
	docker-compose logs -f

# Development setup
dev-setup: install-deps
	@echo "Development setup complete!"

# Full development environment
dev: docker-up
	@echo "Development environment started!"
	@echo "API available at: http://localhost:8080"
	@echo "Use 'make docker-logs' to view logs"
	@echo "Use 'make docker-down' to stop containers" 
