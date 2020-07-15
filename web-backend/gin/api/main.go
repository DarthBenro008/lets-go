package main

import (
	"fmt"
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
	r.Run()
	newsfeeds := newsfeed.New()

	fmt.Println(newsfeeds)
	newsfeeds.Add(newsfeed.Item{"Hello", "How are you ", "Cool"})
	fmt.Println(newsfeeds)

}
