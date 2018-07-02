package routers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Routes the handlers
func Routes(r *gin.Engine) {
	// r.GET("/article/:id", GetArticleHandler)
	// r.GET("/article", GetArticlePreviewsHandler)
	r.POST("/article", InsertArticleHandler)
	// fs := http.FileServer(http.Dir("/usr/share/nick/nick-f/dist/index.html"))
	// http.Handle("/", fs)
	r.GET("/*path", func(c *gin.Context) {
		// /usr/share/nick/nick-f/dist/index.html
		path := c.Param("path")
		fmt.Println(path)

		if strings.HasPrefix(path, "/article/") {
			GetArticleHandler(c)
			return
		}

		if strings.HasPrefix(path, "/article") {
			GetArticlePreviewsHandler(c)
			return
		}

		if path == "" {
			path = "index.html"
		}
		uri := fmt.Sprintf("/usr/share/nick/nick-f/dist/%s", path)
		c.File(uri)
	})

}
