package main

import (
	"gin-practice3/controller"
	"gin-practice3/db"
	"gin-practice3/repository"
	"gin-practice3/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database config
	dbUser := "root"
	dbPassword := "root"
	dbHost := "localhost"
	dbPort := 3303
	dbName := "books_db"

	// Connect to database
	database, err := db.Connect(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	// Create tables
	if err := db.CreateTables(database); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Initialize layers
	bookRepo := repository.NewBookRepository(database)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	// Setup router
	router := gin.Default()

	// Routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running"})
	})

	books := router.Group("/books")
	{
		books.POST("", bookController.CreateBook)
		books.GET("", bookController.GetAllBooks)
		books.GET("/:id", bookController.GetBook)
		books.PUT("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}

	// Start server
	log.Println("Server starting on :8080")
	router.Run(":8080")
}
