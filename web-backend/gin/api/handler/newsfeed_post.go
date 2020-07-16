package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

func PostNewsFeed(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody newsfeed.Item
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		feed.Add(requestBody)
		c.Status(http.StatusNoContent)
	}
}
