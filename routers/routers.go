package routers

import (
	"github.com/gin-gonic/gin"
)

// Routes the handlers
func Routes(r *gin.Engine) {
	r.GET("/blog/article/:id", GetArticleHandler)
	r.GET("/blog/article", GetArticlePreviewsHandler)
	r.POST("/blog/article", AddArticleHandler)

	r.POST("/blog/comment", AddCommentHandler)
	r.POST("/blog/comment/subcomment", AddSubCommentHandler)

	// serve front-end file
	r.GET("/blog/", HTMLFileHandler)
	r.GET("/blog/js/*path", HTMLFileHandler)
	r.GET("/blog/css/*path", HTMLFileHandler)
	r.GET("/blog/favicon.ico", HTMLFileHandler)
}
