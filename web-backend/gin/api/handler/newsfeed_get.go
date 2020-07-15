package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

func GetNewsfeed(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := feed.GetAll()
		c.JSON(http.StatusOK, res)
	}

}
