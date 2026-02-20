package model

import "time"

// Book represents a book in the system
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	ISBN      string    `json:"isbn" binding:"required"`
	Pages     int       `json:"pages" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
	Published time.Time `json:"published" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
