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

func PostNewsFeed(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsFeedPostRequest{}
		c.Bind(&requestBody)
		item := newsfeed.Item{Title: requestBody.Title,
			Post: requestBody.Post,
			Url:  requestBody.Url,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent)
	}
}
