package routers

import (
	"github.com/gin-gonic/gin"
)

// Routes the handlers
func Routes(r *gin.Engine) {
	r.GET("/article/:id", GetArticleHandler)
	r.GET("/article", GetArticlePreviewsHandler)
	r.POST("/article", InsertArticleHandler)
	r.GET("/blog/*path", HTMLFileHandler)
}
