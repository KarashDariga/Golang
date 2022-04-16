package handlers

import (
	"net/http"

	"hw7/models"

	"github.com/gin-gonic/gin"
)

type CreateArticleInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllArticles(context *gin.Context) {
	var articles []models.Article
	models.DB.Find(&articles)

	context.JSON(http.StatusOK, gin.H{"articles": articles})
}

func CreateArticle(context *gin.Context) {

	var input CreateArticleInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&article)

	context.JSON(http.StatusOK, gin.H{"articles": article})
}

func UpdateArticle(context *gin.Context) {
	var article models.Article

	if err := models.DB.Where("id=?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateArticleInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&article).Update(input)
	context.JSON(http.StatusOK, gin.H{"articles": article})
}

func DeleteArticle(context *gin.Context) {
	var article models.Article

	if err := models.DB.Where("id=?", context.Param("id")).First(&article).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&article)
	context.JSON(http.StatusOK, gin.H{"articles": true})
}
