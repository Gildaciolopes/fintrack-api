.PHONY: help build run test clean docker-build docker-up docker-down swagger

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	go build -o bin/api cmd/api/main.go

run: ## Run the application
	go run cmd/api/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

swagger: ## Generate swagger documentation
	swag init -g cmd/api/main.go -o docs

docker-build: ## Build docker image
	docker build -t fintrack-api:latest .

docker-up: ## Start docker containers
	docker-compose up -d

docker-down: ## Stop docker containers
	docker-compose down

docker-logs: ## Show docker logs
	docker-compose logs -f

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...
	gofmt -s -w .

dev: ## Run in development mode with hot reload
	air

.DEFAULT_GOAL := help
