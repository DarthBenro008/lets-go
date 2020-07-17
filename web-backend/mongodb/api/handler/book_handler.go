package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newsfeeder/pkg/book"
	"newsfeeder/pkg/entities"
)

type delRequest struct {
	ID string `json:"id" binding:"required,gte=1"`
}

func AddBook(service book.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody entities.Book
		err := c.Bind(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		res, dberr := service.InsertBook(&requestBody)
		if res != nil {
			c.JSON(http.StatusOK, res)
		}
		c.JSON(http.StatusInternalServerError, dberr)
	}
}

func GetBooks(service book.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		fetched, err := service.FetchBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, fetched)
	}
}

func UpdateBook(service book.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody entities.Book
		err := c.Bind(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		res, dberr := service.UpdateBook(&requestBody)
		if res != nil {
			c.JSON(http.StatusOK, res)
		}
		c.JSON(http.StatusInternalServerError, dberr)
	}
}

func DeleteBook(service book.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody delRequest
		err := c.Bind(&requestBody)
		bookID := requestBody.ID
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		dberr := service.RemoveBook(bookID)
		if dberr != nil {
			c.JSON(http.StatusInternalServerError, dberr)
		}
		c.JSON(http.StatusOK, map[string]string{
			"message": "successful",
		})
	}
}