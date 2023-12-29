package controller

import (
	"bookstore-crud/config"
	"bookstore-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllBooks(ctx *gin.Context) {
	var books []models.BookModel
	config.BookDb.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    books,
	})

}

// Create book
func CreateNewBook(ctx *gin.Context) {
	// validate input
	var input models.CreateBookModel
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Create book
	book := models.BookModel{Title: input.Title, Author: input.Author, PublishedYear: input.PublishedYear, Content: input.Content}
	config.BookDb.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})

}
