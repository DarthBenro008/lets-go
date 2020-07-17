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
//CreateArticle perform C of CRUD
func CreateArticle(article article.Service) gin.HandlerFunc {
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
//ReadArticles performs R of CRUD
func ReadArticles(article article.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,article.FetchAll())
	}
}
//UpdateArticle performs U of CRUD
func UpdateArticle(article article.Service) gin.HandlerFunc {
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
//DeleteArticle performs D of CRUD
func DeleteArticle(article article.Service) gin.HandlerFunc {
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
