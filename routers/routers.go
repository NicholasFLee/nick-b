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
	r.GET("/", func(c *gin.Context) {
		// /usr/share/nick/nick-f/dist/index.html
		fmt.Println("////////")
		c.File("/usr/share/nick/nick-f/dist/index.html")

	})

}
