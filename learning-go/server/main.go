package main

import (
	"github.com/gin-gonic/gin"

	"learning-go/main/controllers"
	"learning-go/main/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
