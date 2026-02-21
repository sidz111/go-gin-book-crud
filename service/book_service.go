package service

import (
	"fmt"
	"gin-practice3/model"
	"gin-practice3/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *model.Book) (*model.Book, error) {
	if err := validateBook(book); err != nil {
		return nil, err
	}

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBook(id int) (*model.Book, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid book id")
	}
	return s.repo.GetByID(id)
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) UpdateBook(id int, book *model.Book) error {
	if id <= 0 {
		return fmt.Errorf("invalid book id")
	}

	existingBook, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Price != 0 {
		existingBook.Price = book.Price
	}

	return s.repo.Update(id, existingBook)
}

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
	if book.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil
}
