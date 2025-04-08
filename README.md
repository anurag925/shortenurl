# URL Shortener Service

A high-performance URL shortening service built with Go, Echo framework, and Bun ORM with SQLite.

## Features

- Create short URLs from long URLs
- Redirect to original URLs using short codes
- Custom alias support
- URL expiration (optional)
- Visit tracking
- RESTful API
- Swagger documentation

## Technologies

- **Go** - Programming language
- **Echo** - High performance web framework
- **Bun** - ORM for Go with SQLite support
- **SQLite** - Embedded database
- **Swagger** - API documentation

## Project Structure
shortenurl/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── api/             # Echo handlers
│   ├── db/
│   │   ├── models/      # Bun ORM models
│   │   ├── migrations/  # SQL migration files
│   │   └── repositories # Database access layer
│   ├── service/         # Business logic
│   └── dto/             # Data transfer objects
├── docs/                # Swagger documentation
└── README.md            # Project documentation

## Installation

1. Clone the repository:
```bash
git clone https://github.com/anurag/shortenurl.git
cd shortenurl
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```
2. Create a .env file in the root directory and configure your environment variables.
3. Run the application:
```bash
go run cmd/main.go
```
http://localhost:8080/swagger/index.html
