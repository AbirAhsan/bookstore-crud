package main

import (
	middlewares "bookstore-crud/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
	server.Run(":8080")
}
