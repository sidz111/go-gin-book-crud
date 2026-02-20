package controller

import (
	"net/http"
	"strconv"

	"gin-practice3/model"
	"gin-practice3/service"

	"github.com/gin-gonic/gin"
)

// BookController handles HTTP requests for books
type BookController struct {
	service *service.BookService
}

// NewBookController creates a new BookController instance
func NewBookController(service *service.BookService) *BookController {
	return &BookController{service: service}
}

// CreateBook handles POST /books
func (c *BookController) CreateBook(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := c.service.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdBook)
}

// GetBook handles GET /books/:id
func (c *BookController) GetBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	book, err := c.service.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// GetAllBooks handles GET /books
func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	ctx.JSON(http.StatusOK, books)
}

// UpdateBook handles PUT /books/:id
func (c *BookController) UpdateBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var book model.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateBook(id, &book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "book updated successfully"})
}

// DeleteBook handles DELETE /books/:id
func (c *BookController) DeleteBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	if err := c.service.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
