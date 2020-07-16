package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

func PutNewsFeed(feeder newsfeed.Updater) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody newsfeed.Item
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := feeder.Update(requestBody)
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
