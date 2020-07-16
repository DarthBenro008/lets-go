package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"
	"github.com/gin-gonic/gin"
)

func PutNewsFeed(feeder newsfeed.Updater) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{Title: requestBody.Title,
			Post: requestBody.Post,
			Url:  requestBody.Url,
		}
		res := feeder.Update(item)
		if res == true {
			c.JSON(http.StatusOK, map[string]string{
				"message": "updated",
			})
		} else {
			c.JSON(http.StatusOK, map[string]string{
				"message": "not found",
			})
		}
	}
}
