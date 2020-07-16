package handlers

import (
	"articlefeeder/pkg/article"
	"articlefeeder/pkg/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type delRequest struct {
	ID int `json:"id" binding:"gte=1"`
}

func CreateArticle(article article.ArticleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody entities.Article
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		res := article.InsertArticle(requestBody)
		c.JSON(http.StatusOK, res)
	}
}

func ReadArticles(article article.ArticleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody entities.Article
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK,article.FetchAll())
	}
}

func UpdateArticle(article article.ArticleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody entities.Article
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK,article.UpdateArticle(requestBody))
	}
	
}

func DeleteArticle(article article.ArticleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody delRequest
		err := c.BindJSON(&requestBody)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		id := requestBody.ID
		res := article.DeleteArticle(id)
		c.JSON(http.StatusOK, res)
	}
}
