package routers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasflee/nick-b/article"
)

// InsertArticleHandler `/article`
// POST: application/x-www-form-urlencoded
// {
// aticleID: string,
// title: string
// categories: [string]
// content: string
// previewContent: string
// }
func InsertArticleHandler(c *gin.Context) {
	fmt.Println("enter insert handler")
	var a article.Article
	if c.ShouldBind(&a) != nil {
		c.JSON(401, gin.H{
			"error": "wrong parameters",
		})
		return
	}
	// now get time
	// RFC3339: "2006-01-02T15:04:05Z07:00"
	a.CreateDate = time.Now().Format(time.RFC3339)
	if article.InsertArticle(a) != nil {
		c.JSON(500, gin.H{
			"error": "insert error",
		})
		return
	}
	c.String(200, "upload success")
}

// GetArticleHandler `/article/:id`
// GET
func GetArticleHandler(c *gin.Context) {
	fmt.Println("enter article handler")
	id := c.Param("id")
	if id == "" {
		c.String(401, "wrong parameter")
		return
	}
	atc, err := article.GetArticle(id)
	if err != nil {
		c.String(500, "get article failed")
		return
	}
	c.JSON(200, atc)
}

// GetArticlePreviewsHandler `/article`
// GET: Query
// page=1&perPage=10
func GetArticlePreviewsHandler(c *gin.Context) {
	fmt.Println("enter previews handler")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.String(401, "'page' is not int")
		return
	}
	perPage, err := strconv.Atoi(c.Query("perPage"))
	if err != nil {
		c.String(401, "'perPage' is not int")
		return
	}
	atcs, err := article.GetArticlePreviews(page, perPage)
	if err != nil {
		c.String(500, "get article failed")
		fmt.Println(err)
		return
	}
	c.JSON(200, atcs)
}
