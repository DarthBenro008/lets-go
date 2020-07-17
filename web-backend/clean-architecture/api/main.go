package main

import (
	"articlefeeder/api/handlers"
	"articlefeeder/pkg/article"

	"github.com/gin-gonic/gin"
)

var (
	articleRepo = article.NewArticleRepo()
	articleSvc  = article.NewService(articleRepo)
)

func main() {
	r := gin.Default()
	r.GET("/articles", handlers.ReadArticles(articleSvc))
	r.POST("/articles", handlers.CreateArticle(articleSvc))
	r.PUT("/articles", handlers.UpdateArticle(articleSvc))
	r.DELETE("/articles", handlers.DeleteArticle(articleSvc))
	_ = r.Run()
}
