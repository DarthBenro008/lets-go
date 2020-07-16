package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

func DeletePost(feeder newsfeed.Deleter) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeed.Item{}
		c.BindJSON(&requestBody)
		item := newsfeed.Item{Title: requestBody.Title,
			Post: requestBody.Post,
			Url:  requestBody.Url,
		}
		res := feeder.Delete(item)
		if res == true {
			c.JSON(http.StatusOK, map[string]string{
				"message": "successful",
			})
		} else {
			c.JSON(http.StatusOK, map[string]string{
				"message": "not found",
			})
		}
	}
}
