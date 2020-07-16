package main

import (
	"newsfeeder/api/handler"
	"newsfeeder/pkg/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.GetNewsfeed(feed))
	r.POST("/newsfeed", handler.PostNewsFeed(feed))
	r.PUT("/newsfeed", handler.PutNewsFeed(feed))
	r.DELETE("/newsfeed", handler.DeletePost(feed))
	r.Run()
}
