.DEFAULT_GOAL := help

.PHONY: help build run dev lint migrate-up migrate-down

help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run in development mode"
	@echo "  make lint         - Run linter"
	@echo "  make migrate-up   - Apply database migrations"
	@echo "  make migrate-down - Rollback last migration"

build:
	mkdir -p bin
	go build -o bin/app ./cmd/api

run:
	go run ./cmd/api

dev:
	go run ./cmd/api

lint:
	golangci-lint run ./...

DB_URL=postgres://postgres:password@localhost:5432/tv_accounts_management?sslmode=disable

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

docker-up:
	docker-compose -f docker/docker-compose.yml up -d

docker-down:
	docker-compose -f docker/docker-compose.yml down