package service

import (
	"fmt"
	"gin-practice3/model"
	"gin-practice3/repository"
	"time"
)

// BookService handles business logic for books
type BookService struct {
	repo *repository.BookRepository
}

// NewBookService creates a new BookService instance
func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook handles book creation
func (s *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	if err := validateBook(book); err != nil {
		return nil, err
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

// GetBook retrieves a book by ID
func (s *BookService) GetBook(id int) (*model.Book, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid book id")
	}
	return s.repo.GetByID(id)
}

// GetAllBooks retrieves all books
func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

// UpdateBook handles book update
func (s *BookService) UpdateBook(id int, book *model.Book) error {
	if id <= 0 {
		return fmt.Errorf("invalid book id")
	}

	existingBook, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Update fields if provided (non-empty/non-zero values)
	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.ISBN != "" {
		existingBook.ISBN = book.ISBN
	}
	if book.Pages != 0 {
		existingBook.Pages = book.Pages
	}
	if book.Price != 0 {
		existingBook.Price = book.Price
	}
	if !book.Published.IsZero() {
		existingBook.Published = book.Published
	}

	existingBook.UpdatedAt = time.Now()

	return s.repo.Update(id, existingBook)
}

// DeleteBook handles book deletion
func (s *BookService) DeleteBook(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid book id")
	}
	return s.repo.Delete(id)
}

// Helper function to validate book
func validateBook(book *model.Book) error {
	if book.Title == "" {
		return fmt.Errorf("title is required")
	}
	if book.Author == "" {
		return fmt.Errorf("author is required")
	}
	if book.ISBN == "" {
		return fmt.Errorf("isbn is required")
	}
	if book.Pages <= 0 {
		return fmt.Errorf("pages must be greater than 0")
	}
	if book.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil
}
