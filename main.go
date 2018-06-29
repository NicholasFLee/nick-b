package main

import (
	"log"

	"github.com/nicholasflee/nick-b/article"
)

func main() {
	a := article.Article{
		ArticleID: "23", Title: "article title 2",
		CreateDate: "2018-09-29T09:12:33.001Z", Categories: []string{"Golang"},
		Content: "**article content", PreviewContent: "article preview content",
	}

	err := article.InsertArticle(a)
	if err != nil {
		log.Fatal(err)
	}

}
