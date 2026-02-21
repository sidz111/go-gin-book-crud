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

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Close()

	if err := db.CreateTables(database); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	bookRepo := repository.NewBookRepository(database)
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	router := gin.Default()

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

	log.Println("Server starting on :8080")
	router.Run(":8080")
}
