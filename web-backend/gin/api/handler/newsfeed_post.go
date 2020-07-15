package handler

import (
	"net/http"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsFeedPostRequest struct {
	Title string `json: "title`
	Post  string `json: "post"`
	Url   string `json: "url`
}

func PostNewsFeed(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)
		feed.Add(newsfeed.Item{Title: requestBody.Title,
			Post: requestBody.Post,
			Url:  requestBody.Url,
		})
		c.Status(http.StatusNoContent)
	}
}
