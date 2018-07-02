package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/nicholasflee/nick-b/article"
	"github.com/nicholasflee/nick-b/routers"
)

func main() {
	a := article.Article{
		ArticleID: "23", Title: "article title 2",
		CreateDate: "2018-09-29T09:12:33.001Z", Categories: []string{"Golang"},
		Content: "**article content", PreviewContent: "article preview content",
	}

	fmt.Println(a)

	r := gin.Default()
	routers.Routes(r)
	// err := article.InsertArticle(a)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r.Run()
}
