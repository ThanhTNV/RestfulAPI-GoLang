# RestfulAPI-GoLang

A RESTful API server built with Go, Gin framework, and PostgreSQL, following a clean 3-tier architecture with dependency injection, running in Docker Compose.

## Features

- **3-Tier Architecture**: Presentation (API), Service (Business Logic), and Repository (Data Access) layers
- **Dependency Injection**: Modular and testable design
- **Router-Middleware-Controller Pattern**: Clean separation of concerns
- RESTful API with CRUD operations
- Gin web framework with custom middleware
- PostgreSQL database with GORM ORM
- Docker Compose for easy deployment
- Health check endpoint
- Error handling middleware

## Architecture

This project follows a clean 3-tier architecture:

### Presentation Layer (API Layer)
- **Controllers**: Handle HTTP requests and responses
- **Middleware**: Error handling, logging, and request validation
- **Routes**: Define API endpoints

### Service Layer
- Contains business logic
- Orchestrates between repositories
- Independent of transport layer (HTTP, CLI, etc.)

### Repository Layer
- Handles data persistence
- Maps database rows to Go structs
- Abstracts database operations

## Prerequisites

- Docker
- Docker Compose

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/ThanhTNV/RestfulAPI-GoLang.git
cd RestfulAPI-GoLang
```

2. Start the application with Docker Compose:
```bash
docker compose up --build
```

The API server will be available at `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Check if the server is running

### Users API
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get a specific user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

## API Usage Examples

### Create a user
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

### Get all users
```bash
curl http://localhost:8080/api/v1/users
```

### Get a specific user
```bash
curl http://localhost:8080/api/v1/users/1
```

### Update a user
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Development

### Running locally without Docker

1. Install Go (1.24 or later)

2. Install PostgreSQL and create a database:
```bash
createdb restapi
```

3. Set environment variables:
```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=restapi
export DB_PORT=5432
export PORT=8080
```

4. Run the application:
```bash
go run cmd/api/main.go
```

## Environment Variables

- `DB_HOST` - Database host (default: localhost)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: restapi)
- `DB_PORT` - Database port (default: 5432)
- `PORT` - API server port (default: 8080)

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point with DI setup
├── internal/
│   ├── api/
│   │   ├── controllers/         # HTTP request handlers
│   │   ├── middleware/          # Error handling, logging, etc.
│   │   └── routes/              # Route definitions
│   ├── service/                 # Business logic layer
│   ├── repository/              # Data access layer
│   └── models/                  # Domain entities
├── pkg/
│   ├── config/                  # Configuration management
│   └── database/                # Database connection utilities
├── Dockerfile                   # Docker configuration
├── docker-compose.yml           # Docker Compose configuration
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies
└── README.md                    # This file
```

## Design Patterns

### Dependency Injection
The application uses constructor-based dependency injection. Dependencies flow from `main.go`:
```
main.go → Repository → Service → Controller → Routes
```

### Middleware Pattern
Custom middleware intercepts requests for:
- Error handling: Centralized error response formatting
- Recovery: Panic recovery with graceful error messages
- Logging: Request/response logging (can be extended)

### Repository Pattern
Abstract database operations behind interfaces, making it easy to:
- Swap implementations
- Mock for testing
- Change databases without affecting business logic

## License

MIT
