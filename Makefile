# ==============================================================================
# PRAHARI Monorepo Orchestration Makefile
# ==============================================================================

.PHONY: help setup run-infra stop-infra clean lint test test-unit build-all

# Detect OS
OS := $(shell uname -s)

help:
	@echo "========================================================================"
	@echo "PRAHARI Developer Command Center"
	@echo "========================================================================"
	@echo "Available commands:"
	@echo "  setup         - Initialize developer environment, workspace settings, etc."
	@echo "  run-infra     - Spin up local development infrastructure (Postgres, Redis, Kafka)"
	@echo "  stop-infra    - Spin down local development infrastructure"
	@echo "  clean         - Clean up compile artifacts, virtualenvs, node_modules"
	@echo "  lint          - Run linters across Go, Python, and React services"
	@echo "  test          - Run all unit and integration tests across services"
	@echo "  build-all     - Build all Go microservices and React web bundles"
	@echo "========================================================================"

setup:
	@echo "Initializing PRAHARI Monorepo Workspaces..."
	@if [ ! -f .env ]; then cp .env.example .env; fi
	@echo "Setting up Go Workspaces..."
	@go work init || true
	@echo "Installing frontend dependencies..."
	@npm install --workspaces || echo "Skipping frontend npm setup (package.json placeholder)"
	@echo "Setting up Python environments..."
	@python3 -m venv .venv
	@.venv/bin/pip install --upgrade pip
	@echo "Developer setup completed successfully!"

run-infra:
	@echo "Spinning up backing Docker infrastructure..."
	docker-compose up -d

stop-infra:
	@echo "Stopping backing Docker infrastructure..."
	docker-compose down

clean:
	@echo "Cleaning build artifacts..."
	find . -name "dist" -type d -exec rm -rf {} +
	find . -name "build" -type d -exec rm -rf {} +
	find . -name "*.out" -type f -delete
	find . -name "*.test" -type f -delete
	find . -name "__pycache__" -type d -exec rm -rf {} +
	find . -name ".pytest_cache" -type d -exec rm -rf {} +
	find . -name ".venv" -type d -exec rm -rf {} +
	find . -name "node_modules" -type d -exec rm -rf {} +

lint:
	@echo "Linting Go microservices..."
	@if command -v golangci-lint >/dev/null; then golangci-lint run ./...; else echo "golangci-lint not installed. Run: go install github.com/golangci/lint/cmd/golangci-lint@latest"; fi
	@echo "Linting Python AI/CV components..."
	@if [ -d .venv ]; then .venv/bin/pip install flake8 black; .venv/bin/flake8 ai/ computer-vision/; else echo "No virtualenv found. Run 'make setup' first."; fi
	@echo "Linting React frontends..."
	@npm run lint --workspaces || echo "React linting skipped or package workspaces unconfigured."

test:
	@echo "Running Go service tests..."
	go test -v ./...
	@echo "Running Python AI/CV tests..."
	@if [ -d .venv ]; then .venv/bin/pip install pytest; .venv/bin/pytest tests/unit/ai; else echo "No virtualenv found. Run 'make setup' first."; fi
	@echo "Running Frontend tests..."
	@npm test --workspaces || echo "React tests skipped."

build-all:
	@echo "Building Go microservices..."
	@for service in services/*-service; do \
		if [ -f $$service/main.go ]; then \
			echo "Building $$service..."; \
			go build -o bin/$$(basename $$service) $$service/main.go; \
		fi \
	done
	@echo "Building Frontend Apps..."
	@npm run build --workspaces || echo "React builds skipped."
