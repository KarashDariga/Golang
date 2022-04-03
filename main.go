package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []article{
	{
		ID:      "1",
		Title:   "Title 1",
		Content: "Body 1",
	},
	{
		ID:      "2",
		Title:   "Title 2",
		Content: "Body 2",
	},
	{
		ID:      "3",
		Title:   "Title 3",
		Content: "Body 3",
	},
}

func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

func postArticles(c *gin.Context) {
	var article article

	if err := c.BindJSON(&article); err != nil {
		return
	}

	articles = append(articles, article)
	c.IndentedJSON(http.StatusCreated, article)
}

func getArticleByID(c *gin.Context) {
	id := c.Param("id")

	for _, article := range articles {
		if article.ID == id {
			c.IndentedJSON(http.StatusOK, article)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

func deleteArticleByID(c *gin.Context) {
	id := c.Param("id")

	for i, article := range articles {
		if article.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

func main() {
	router := gin.Default()
	router.GET("/articles", getArticles)
	router.POST("/postArticles", postArticles)
	router.GET("/articles/:id", getArticleByID)
	router.DELETE("/deleteArticle/:id", deleteArticleByID)

	router.Run("localhost:8080")
}
