# ---------- Stage 1: Build ----------
FROM golang:1.24 AS builder

WORKDIR /app

# Сначала копируем go.mod для кэширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-profiler ./cmd/server

# ---------- Stage 2: Runtime ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-profiler .

EXPOSE 8080

CMD ["./api-profiler"]