# Forum Project with Admin Panel

This project implements a forum system with an admin panel using microservices architecture.

## Architecture

The project consists of two main microservices:
1. Auth Service - handles user authentication and management
2. Forum Service - handles forum posts, comments, and chat functionality

## Technologies Used

- Go
- PostgreSQL
- gRPC for inter-service communication
- WebSocket for real-time chat
- golang-migrate for database migrations
- Zap for logging
- Swagger for API documentation
- Docker for containerization

## Project Structure

```
.
├── auth-service/           # Authentication microservice
│   ├── cmd/               # Application entry points
│   ├── internal/          # Private application code
│   ├── pkg/              # Public library code
│   └── migrations/       # Database migrations
│
├── forum-service/         # Forum microservice
│   ├── cmd/              # Application entry points
│   ├── internal/         # Private application code
│   ├── pkg/             # Public library code
│   └── migrations/      # Database migrations
│
└── proto/                # Protocol buffer definitions
```

## Setup Instructions

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up PostgreSQL database

3. Run migrations:
   ```bash
   make migrate-up
   ```

4. Start services:
   ```bash
   make run-services
   ```

## API Documentation

API documentation is available at `/swagger/index.html` when the services are running.

## Testing

Run tests:
```bash
make test
```

## Features

- User authentication and authorization
- Forum posts and comments
- Real-time chat with WebSocket
- Admin panel
- Message auto-deletion
- Comprehensive logging
- API documentation with Swagger 