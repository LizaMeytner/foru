# Сервис форума

Это компонент форума в микросервисной архитектуре приложения. Он обрабатывает посты, комментарии и сообщения чата.

## Требования

- Go 1.21 или новее
- PostgreSQL 12 или новее
- Protocol Buffers compiler (protoc)
- Go плагины для protoc
- Docker и Docker Compose (для запуска в контейнерах)

## Конфигурация

Сервис можно настроить с помощью переменных окружения или файла `.env`. Вот доступные параметры конфигурации:

```env
GRPC_PORT=50051

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=forum
POSTGRES_SSLMODE=disable
```

## Запуск с помощью Docker

1. Сборка и запуск контейнеров:
```bash
docker-compose up --build
```

2. Для запуска в фоновом режиме:
```bash
docker-compose up -d
```

3. Для остановки:
```bash
docker-compose down
```

## Запуск без Docker

### Настройка базы данных

1. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE forum;
```

2. Запустите миграции:
```bash
go run cmd/migrate/main.go
```

### Запуск сервиса

1. Установите зависимости:
```bash
go mod download
```

2. Сгенерируйте код из protobuf:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/forum.proto
```

3. Запустите сервис:
```bash
go run cmd/api/main.go
```

## API Documentation

Сервис предоставляет следующие gRPC эндпоинты:

### Посты
- CreatePost - Создать пост
- GetPost - Получить пост
- ListPosts - Список постов
- UpdatePost - Обновить пост
- DeletePost - Удалить пост

### Комментарии
- CreateComment - Создать комментарий
- GetComment - Получить комментарий
- ListComments - Список комментариев
- UpdateComment - Обновить комментарий
- DeleteComment - Удалить комментарий

### Чат
- CreateChatMessage - Создать сообщение чата
- GetChatMessage - Получить сообщение чата
- ListChatMessages - Список сообщений чата
- DeleteOldChatMessages - Удалить старые сообщения чата

Для подробной документации API обратитесь к файлу `proto/forum.proto`. 