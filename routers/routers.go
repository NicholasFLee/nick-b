package routers

import (
	"github.com/gin-gonic/gin"
)

// Routes the handlers
func Routes(r *gin.Engine) {
	r.GET("/blog/article/:id", GetArticleHandler)
	r.GET("/blog/article", GetArticlePreviewsHandler)
	r.POST("/blog/article", InsertArticleHandler)
	r.GET("/blog/*path", HTMLFileHandler)
}
