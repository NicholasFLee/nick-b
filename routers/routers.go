package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Routes the handlers
func Routes(r *gin.Engine) {
	r.GET("/article/:id", GetArticleHandler)
	r.GET("/article", GetArticlePreviewsHandler)
	r.POST("/article", InsertArticleHandler)
	// fs := http.FileServer(http.Dir("/usr/share/nick/nick-f/dist/index.html"))
	// http.Handle("/", fs)
	r.GET("/blog/*path", func(c *gin.Context) {
		// /usr/share/nick/nick-f/dist/index.html
		path := c.Param("path")
		if path == "" {
			path = "index.html"
		}
		uri := fmt.Sprintf("/usr/share/nick/nick-f/dist/%s", path)
		c.File(uri)
	})

}
