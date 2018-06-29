package routers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasflee/nick-b/article"
)

// InsertArticleHandler `/article`
// POST: application/x-www-form-urlencoded
// {
// id: string,
// title: string
// categories: [string]
// content: string
// previewContent: string
// }
func InsertArticleHandler(c *gin.Context) {
	var a article.Article
	if c.ShouldBind(&a) != nil {
		c.JSON(401, gin.H{
			"error": "wrong parameters",
		})
	}
	// now get time
	// RFC3339: "2006-01-02T15:04:05Z07:00"
	a.CreateDate = time.Now().Format(time.RFC3339)
	if article.InsertArticle(a) != nil {
		c.JSON(500, gin.H{
			"error": "insert error",
		})
	}
	c.String(200, "upload success")
}

// GetArticleHandler `/article/:id`
// GET
func GetArticleHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(401, "wrong parameter")
	}
	atc, err := article.GetArticle(id)
	if err != nil {
		c.String(500, "get article failed")
	}
	c.JSON(200, atc)
}

// GetArticlePreviewsHandler `/article`
// GET: Query
// page=1&perPage=10
func GetArticlePreviewsHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.String(401, "'page' is not int")
	}
	perPage, err := strconv.Atoi(c.Query("perPage"))
	if err != nil {
		c.String(401, "'perPage' is not int")
	}
	atcs, err := article.GetArticlePreviews(page, perPage)
	if err != nil {
		c.String(500, "get article failed")
	}
	c.JSON(200, atcs)
}
