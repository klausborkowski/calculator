.PHONY: test test-report docker-up docker-down docker-build generate-doc dev-up dev-down start

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

docker-build:
	docker compose build
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
	swag init -d internal/api -g calculator_api.go -g package.go -g server_side_rendering.go
