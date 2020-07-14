package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json: "desc"`
	Content string `json:"content"`
}

type Articles []Article

func showAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Print("All Articel has been hit")
	articles := Articles{
		Article{Title: "This is title", Desc: "This is desc", Content: "This is sicc content"},
	}
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go has been hit")
}

func requestHandler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", showAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	requestHandler()

}
