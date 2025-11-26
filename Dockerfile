FROM golang:1.22.6 AS builder

WORKDIR /app

# Копируем go mod файлы
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o packager ./cmd/server/main.go

FROM alpine:latest

# Устанавливаем необходимые пакеты
RUN apk --no-cache add ca-certificates wget

WORKDIR /root/

# Копируем бинарник
COPY --from=builder /app/packager .

# Копируем .env файл если существует (опционально)
COPY --from=builder /app/.env* ./

EXPOSE 8080

CMD ["./packager"]