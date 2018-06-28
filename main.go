package main

import (
	"log"

	"github.com/nicholasflee/nick-b/article"
)

func main() {
	a := article.Article{
		ID: "rew", Title: "title", Subtitle: "subtitle",
		CreateDate: "date", Categories: []string{"12"},
		Content: "", PreviewContent: "",
	}

	err := article.InsertArticle(a)
	if err != nil {
		log.Fatal(err)
	}
}
