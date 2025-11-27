.PHONY: test test-report docker-up docker-down docker-build docker-build-backend docker-build-frontend build-local build-backend build-frontend generate-doc dev-up dev-down start start-local

BIN_DIR ?= bin
BACKEND_BINARY ?= $(BIN_DIR)/packager
BACKEND_IMAGE ?= calculator-backend
FRONTEND_IMAGE ?= calculator-frontend

# Build commands
build-backend:
	@echo "Building backend binary into $(BACKEND_BINARY)..."
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 go build -o $(BACKEND_BINARY) ./cmd/server/main.go

build-frontend:
	@echo "Building frontend assets..."
	@cd frontend && npm install && npm run build

build-local: build-backend build-frontend

# Docker build commands
docker-build: docker-build-backend docker-build-frontend

docker-build-backend:
	@echo "Building backend Docker image: $(BACKEND_IMAGE)"
	@docker build -t $(BACKEND_IMAGE) -f Dockerfile .

docker-build-frontend:
	@echo "Building frontend Docker image: $(FRONTEND_IMAGE)"
	@docker build -t $(FRONTEND_IMAGE) -f frontend/Dockerfile frontend

# Docker commands
start: docker-up

docker-up:
	@echo "Starting package service set-up..."
	@docker compose up --build -d
	@echo "Waiting for services to start..."
	@sleep 5
	@docker compose ps
	@echo ""
	@echo "Service UI started on http://localhost:3000"
	@echo "Service Backend started on http://localhost:8080"
	@echo "Service Swagger available on http://localhost:8080/swagger/"
	@echo ""
	@echo "To view logs, run: docker compose logs -f"
	@echo "To stop services, run: make docker-down"

docker-down:
	docker compose down

test:
	go test ./...

test-report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html



# Local development (run without Docker)
dev-up:
	@echo "Starting backend and frontend..."
	@echo "Backend: http://localhost:8080"
	@echo "Frontend: http://localhost:3000"
	@trap 'kill 0' EXIT; \
	go run ./cmd/server/main.go & \
	cd frontend && npm run dev & \
	wait

dev-down:
	@echo "Stopping processes..."
	@pkill -f "go run ./cmd/server/main.go" || true
	@pkill -f "vite" || true

start-local:
	@bash -lc '\
		set -a; \
		if [ -f .env ]; then \
			echo "Loading environment from .env"; \
			. ./.env; \
		fi; \
		set +a; \
		echo "Starting PostgreSQL container..."; \
		docker compose up postgres -d; \
		trap "echo \"Stopping local services...\"; kill 0; docker compose stop postgres >/dev/null 2>&1 || true" EXIT; \
		: "$${DB_HOST:=localhost}"; \
		: "$${DB_PORT:=5432}"; \
		: "$${DB_USER:=calculator}"; \
		: "$${DB_PASSWORD:=calculator}"; \
		: "$${DB_NAME:=calculator}"; \
		: "$${PORT:=8080}"; \
		: "$${VITE_API_BASE_URL:=http://localhost:8080}"; \
		export DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME PORT VITE_API_BASE_URL; \
		echo "Starting backend on $$PORT and frontend dev server (VITE_API_BASE_URL=$$VITE_API_BASE_URL)..."; \
		go run ./cmd/server/main.go & \
		( cd frontend && npm install >/dev/null 2>&1 && npm run dev ) & \
		wait \
	'

# Old command for backward compatibility
IMAGE_NAME=packager
CONTAINER_NAME=packager
PORT=8080
docker-up-old:
	docker build -t $(IMAGE_NAME) . && \
	docker stop $(CONTAINER_NAME) 2>/dev/null || true && \
	docker rm $(CONTAINER_NAME) 2>/dev/null || true && \
	docker run -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

generate-doc:
	swag init -d internal/api -g calcalator_api.go -g package.go
