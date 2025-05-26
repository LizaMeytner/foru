# Этап сборки
FROM golang:1.21-alpine AS builder

# Установка необходимых зависимостей
RUN apk add --no-cache git protobuf-dev

# Установка protoc-gen-go и protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Установка рабочей директории
WORKDIR /app

# Копирование файлов проекта
COPY . .

# Генерация protobuf кода
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/forum.proto

# Сборка приложения
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/forum-service ./cmd/api

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копирование бинарного файла из этапа сборки
COPY --from=builder /app/forum-service .
COPY --from=builder /app/migrations ./migrations

# Установка необходимых зависимостей
RUN apk add --no-cache ca-certificates

# Запуск приложения
CMD ["./forum-service"] 