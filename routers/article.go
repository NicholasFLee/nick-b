package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasflee/nick-b/article"
)

// InsertArticleHandler `/article`
// POST: application/x-www-form-urlencoded
// {
// title: string
// categories: [string]
// content: string
// previewContent: string
// }
func InsertArticleHandler(c *gin.Context) {
	fmt.Println("enter insert handler")
	var a article.Article
	if err := c.ShouldBind(&a); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"error": "wrong parameters",
		})
		return
	}
	// set CreateDate
	// RFC3339: "2006-01-02T15:04:05Z07:00"
	a.CreateDate = time.Now().Format(time.RFC3339)
	// set ArticleID -> md5 title
	h := md5.New()
	io.WriteString(h, a.Title)
	atcID := hex.EncodeToString(h.Sum(nil))
	a.ArticleID = atcID
	fmt.Println(atcID)
	if err := article.InsertArticle(a); err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
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
