# RestfulAPI-GoLang

A RESTful API server built with Go, Gin framework, and PostgreSQL, running in Docker Compose.

## Features

- RESTful API with CRUD operations
- Gin web framework
- PostgreSQL database
- GORM ORM
- Docker Compose for easy deployment
- Health check endpoint

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
go run main.go
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
├── main.go              # Main application file
├── Dockerfile           # Docker configuration
├── docker-compose.yml   # Docker Compose configuration
├── go.mod              # Go module file
├── go.sum              # Go dependencies
└── README.md           # This file
```

## License

MIT
