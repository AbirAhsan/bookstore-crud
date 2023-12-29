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
