# Go Chat Server Makefile

# Variables
APP_NAME := go-chat-server
CMD_DIR := cmd/server
BUILD_DIR := build
BINARY := $(BUILD_DIR)/$(APP_NAME)

# Go related variables
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod
GOFMT := $(GOCMD) fmt
GOVET := $(GOCMD) vet

# Build flags
LDFLAGS := -ldflags "-w -s"
BUILD_FLAGS := -v $(LDFLAGS)

# Default target
.PHONY: all
all: clean build

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY) ./$(CMD_DIR)
	@echo "Build completed: $(BINARY)"

# Run the application
.PHONY: run
run:
	@echo "Starting $(APP_NAME)..."
	$(GOCMD) run ./$(CMD_DIR)

# Run with development mode
.PHONY: dev
dev:
	@echo "Starting $(APP_NAME) in development mode..."
	DEV_MODE=true $(GOCMD) run ./$(CMD_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	$(GOVET) ./...

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Tidy dependencies
.PHONY: tidy
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Install the application
.PHONY: install
install:
	@echo "Installing $(APP_NAME)..."
	$(GOCMD) install ./$(CMD_DIR)

# Check code quality
.PHONY: check
check: fmt vet lint test

# Build for different platforms
.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 ./$(CMD_DIR)

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe ./$(CMD_DIR)

.PHONY: build-mac
build-mac:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 ./$(CMD_DIR)

# Build for all platforms
.PHONY: build-all
build-all: build-linux build-windows build-mac

# Docker related targets
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --name $(APP_NAME) $(APP_NAME):latest

.PHONY: docker-stop
docker-stop:
	@echo "Stopping Docker container..."
	@docker stop $(APP_NAME) 2>/dev/null || true
	@docker rm $(APP_NAME) 2>/dev/null || true

.PHONY: docker-clean
docker-clean:
	@echo "Cleaning Docker images..."
	@docker rmi $(APP_NAME):latest 2>/dev/null || true
	@docker system prune -f

# Docker Compose targets
.PHONY: docker-up
docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

.PHONY: docker-down
docker-down:
	@echo "Stopping services with Docker Compose..."
	docker-compose down

.PHONY: docker-logs
docker-logs:
	@echo "Showing Docker Compose logs..."
	docker-compose logs -f

.PHONY: docker-restart
docker-restart: docker-down docker-up

# Development environment
.PHONY: docker-dev-up
docker-dev-up:
	@echo "Starting development environment..."
	docker-compose -f docker/docker-compose.dev.yml up -d

.PHONY: docker-dev-down
docker-dev-down:
	@echo "Stopping development environment..."
	docker-compose -f docker/docker-compose.dev.yml down

.PHONY: docker-dev-logs
docker-dev-logs:
	@echo "Showing development logs..."
	docker-compose -f docker/docker-compose.dev.yml logs -f

# Production environment
.PHONY: docker-prod-up
docker-prod-up:
	@echo "Starting production environment..."
	docker-compose -f docker/docker-compose.prod.yml up -d

.PHONY: docker-prod-down
docker-prod-down:
	@echo "Stopping production environment..."
	docker-compose -f docker/docker-compose.prod.yml down

.PHONY: docker-prod-logs
docker-prod-logs:
	@echo "Showing production logs..."
	docker-compose -f docker/docker-compose.prod.yml logs -f

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          Build the application"
	@echo "  run            Run the application"
	@echo "  dev            Run in development mode"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage report"
	@echo "  fmt            Format code"
	@echo "  vet            Vet code"
	@echo "  lint           Lint code (requires golangci-lint)"
	@echo "  tidy           Tidy dependencies"
	@echo "  deps           Download dependencies"
	@echo "  clean          Clean build artifacts"
	@echo "  install        Install the application"
	@echo "  check          Run fmt, vet, lint, and test"
	@echo "  build-linux    Build for Linux"
	@echo "  build-windows  Build for Windows"
	@echo "  build-mac      Build for macOS"
	@echo "  build-all      Build for all platforms"
	@echo ""
	@echo "Docker Commands:"
	@echo "  docker-build   Build Docker image"
	@echo "  docker-run     Run Docker container"
	@echo "  docker-stop    Stop and remove Docker container"
	@echo "  docker-clean   Clean Docker images and system"
	@echo "  docker-up      Start services with docker-compose"
	@echo "  docker-down    Stop services with docker-compose"
	@echo "  docker-logs    Show docker-compose logs"
	@echo "  docker-restart Restart docker-compose services"
	@echo ""
	@echo "Environment-specific Commands:"
	@echo "  docker-dev-up     Start development environment"
	@echo "  docker-dev-down   Stop development environment"
	@echo "  docker-dev-logs   Show development logs"
	@echo "  docker-prod-up    Start production environment"
	@echo "  docker-prod-down  Stop production environment"
	@echo "  docker-prod-logs  Show production logs"
	@echo ""
	@echo "  help           Show this help message"