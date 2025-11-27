FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download



RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o packager ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/packager .
COPY --from=builder /app/.env .


EXPOSE 8080

CMD ["./packager"]