package main

import (
	"bookstore-crud/config"
	"bookstore-crud/controller"
	middlewares "bookstore-crud/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	config.ConnectBookDatabase()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
	server.GET("/books", controller.FindAllBooks)
	server.POST("/createBook", controller.CreateNewBook)
	server.GET("/book/:id", controller.FindBookById)
	server.PATCH("/updateBook/:id", controller.UpdateBookById)
	server.Run(":8080")
}
