package routers

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nicholasflee/nick-b/comment"
)

// GetCommentsHandler get comments by article id
// func GetCommentsHandler(c *gin.Context) {

// }

// AddCommentHandler add a comment to a article
/**
POST: {
	articleID string
	content string
	authorName string
}
*/
func AddCommentHandler(c *gin.Context) {
	cmt := comment.Comment{}
	// content authorName
	err := c.ShouldBind(&cmt)
	if err != nil {
		c.String(400, "wrong param")
	}
	// aticleID
	atcID := c.PostForm("articleID")
	if atcID == "" {
		c.String(400, "wrong param")
	}
	// date
	cmt.CreateDate = time.Now().Format(time.RFC3339)
	// id
	cmt.CommentID = md5String(cmt.CreateDate + cmt.Content)

	comment.AddComment(atcID, cmt)
}

// AddSubCommentHandler add a sub-comment to a comment
func AddSubCommentHandler(c *gin.Context) {

}

func md5String(original string) string {
	h := md5.New()
	io.WriteString(h, original)
	return hex.EncodeToString(h.Sum(nil))
}
