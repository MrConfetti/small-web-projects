package main

import (
	"github.com/gin-gonic/gin"

	"learning-go/main/controllers"
	"learning-go/main/models"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(cors())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run(":8080")
}
