package main

import (
	"hw7/handlers"
	"hw7/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/articles", handlers.GetAllArticles)
	route.POST("/createArticle", handlers.CreateArticle)
	route.PUT("/update/:id", handlers.UpdateArticle)
	route.DELETE("/delete/:id", handlers.DeleteArticle)

	route.Run()
}
