.PHONY: test test-report docker-up generate-doc

test:
	go test ./...

test-report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

#docker-up
IMAGE_NAME=packager
CONTAINER_NAME=packager
PORT=8080
docker-up:
	docker build -t $(IMAGE_NAME) . && \
	docker stop $(CONTAINER_NAME) 2>/dev/null || true && \
	docker rm $(CONTAINER_NAME) 2>/dev/null || true && \
	docker run -p $(PORT):8080 --name $(CONTAINER_NAME) $(IMAGE_NAME)

generate-doc:
	swag init -d internal/api -g calculator_api.go -g package.go -g server_side_rendering.go
