# Books Management API - Simple Gin + MySQL CRUD

A simple CRUD RESTful API built with Go's Gin framework and MySQL for managing books.

## Project Structure

```
gin-practice3/
├── model/                  # Data models
│   └── book.go
├── db/                     # Database connection
│   └── db.go
├── repository/             # Data access layer
│   └── book_repository.go
├── service/                # Business logic layer
│   └── book_service.go
├── controller/             # HTTP handlers
│   └── book_controller.go
├── main.go                 # Application entry point
├── go.mod
├── go.sum
└── README.md
```

## Architecture Layers

1. **Model** - Book data structures
2. **DB** - Database connection and table creation
3. **Repository** - CRUD database operations (no GORM)
4. **Service** - Business logic and validation
5. **Controller** - HTTP request/response handling

## Prerequisites

- Go 1.21+
- MySQL 5.7+
- MySQL running on localhost:3303

## Setup Instructions

### 1. Create MySQL Database

```bash
mysql -u root -p
CREATE DATABASE books_db;
```

### 2. Install Dependencies

```bash
cd gin-practice3
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### Health Check

```bash
GET /health
```

### Create Book

```bash
POST /books
Content-Type: application/json

{
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "isbn": "978-0134190440",
  "pages": 400,
  "price": 49.99,
  "published": "2015-10-26T00:00:00Z"
}
```

### Get All Books

```bash
GET /books
```

### Get Book by ID

```bash
GET /books/1
```

### Update Book

```bash
PUT /books/1
Content-Type: application/json

{
  "title": "Updated Title",
  "price": 59.99
}
```

### Delete Book

```bash
DELETE /books/1
```

## Database Schema

The books table is created automatically:

```sql
CREATE TABLE IF NOT EXISTS books (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    pages INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    published DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## Database Configuration

Edit `main.go` to change database credentials:

```go
dbUser := "root"
dbPassword := "root"
dbHost := "localhost"
dbPort := 3303
dbName := "books_db"
```

## Error Responses

All errors return appropriate HTTP status codes:

- `400 Bad Request` - Invalid input
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Testing with cURL

```bash
# Create a book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Go Book","author":"John","isbn":"123","pages":300,"price":29.99,"published":"2020-01-01T00:00:00Z"}'

# Get all books
curl http://localhost:8080/books

# Get book with ID 1
curl http://localhost:8080/books/1

# Update book
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Go Book","price":39.99}'

# Delete book
curl -X DELETE http://localhost:8080/books/1

# Health check
curl http://localhost:8080/health
```
