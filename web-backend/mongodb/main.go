package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"newsfeeder/api/handler"
	"newsfeeder/pkg/book"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("server.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("DB connection success!!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	r := gin.Default()
	r.GET("/books", handler.GetBooks(bookService))
	r.POST("/books", handler.AddBook(bookService))
	r.PUT("/books", handler.UpdateBook(bookService))
	r.DELETE("/books", handler.DeleteBook(bookService))
	_ = r.Run()
}

func DatabaseConnection() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_KEY")))
	if err != nil {
		return nil, err
	}
	db := client.Database("books")
	return db, nil
}
