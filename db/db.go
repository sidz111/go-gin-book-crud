package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect creates a MySQL database connection
func Connect() (*sql.DB, error) {

	dbUser := "root"
	dbPass := "root"
	host := "localhost"
	port := 3303
	dbName := "books_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully!")
	return db, nil
}

// CreateTables creates the books table
func CreateTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Tables created successfully!")
	return nil
}
