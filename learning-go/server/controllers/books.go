package controllers

import (
	"github.com/gin-gonic/gin"
	"learning-go/main/models"
	"net/http"
)

// CreateBookInput struct to validate input
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pages  uint   `json:"pages" binding:"required"`
}

// POST /books
// Endpoint to create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author, Pages: input.Pages}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /books
// Endpoint to get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
