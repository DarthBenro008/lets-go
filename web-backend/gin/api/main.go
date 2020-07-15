package main

import (
	"fmt"
	"newsfeeder/pkg/newsfeed"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	// r.Run()
	newsfeeds := newsfeed.New()

	fmt.Println(newsfeeds)
	newsfeeds.Add(newsfeed.Item{"Hello", "How are you ", "Cool"})
	fmt.Println(newsfeeds)
}
