package repository

import (
	"database/sql"
	"fmt"
	"gin-practice3/model"
)

// BookRepository handles all database operations for books
type BookRepository struct {
	db *sql.DB
}

// NewBookRepository creates a new BookRepository instance
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// Create inserts a new book into the database
func (r *BookRepository) Create(book *model.Book) error {
	query := `INSERT INTO books (title, author, isbn, pages, price, published) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, book.Title, book.Author, book.ISBN, book.Pages, book.Price, book.Published)
	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)
	return nil
}

// GetByID retrieves a book by its ID
func (r *BookRepository) GetByID(id int) (*model.Book, error) {
	book := &model.Book{}
	query := `SELECT id, title, author, isbn, pages, price, published, created_at, updated_at FROM books WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages,
		&book.Price, &book.Published, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to get book: %w", err)
	}

	return book, nil
}

// GetAll retrieves all books from the database
func (r *BookRepository) GetAll() ([]model.Book, error) {
	var books []model.Book
	query := `SELECT id, title, author, isbn, pages, price, published, created_at, updated_at FROM books ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(
			&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages,
			&book.Price, &book.Published, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		books = append(books, book)
	}

	return books, nil
}

// Update updates an existing book
func (r *BookRepository) Update(id int, book *model.Book) error {
	query := `UPDATE books SET title = ?, author = ?, isbn = ?, pages = ?, price = ?, published = ? WHERE id = ?`

	result, err := r.db.Exec(query, book.Title, book.Author, book.ISBN, book.Pages, book.Price, book.Published, id)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

// Delete deletes a book by its ID
func (r *BookRepository) Delete(id int) error {
	query := `DELETE FROM books WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}
