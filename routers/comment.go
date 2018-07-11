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
		return
	}
	// aticleID
	atcID := c.PostForm("articleID")
	if atcID == "" {
		c.String(400, "wrong articleID")
		return
	}
	// date
	cmt.CreateDate = time.Now().Format(time.RFC3339)
	// id
	cmt.CommentID = md5String(cmt.CreateDate + cmt.Content)

	err = comment.AddComment(atcID, cmt)
	if err != nil {
		c.String(500, "add comment failed")
		return
	}
	c.String(200, "Upload succeeded")
}

// AddSubCommentHandler add a sub-comment to a comment
/**
POST: {
	commentID string
	content string
	authorName string
}
*/
func AddSubCommentHandler(c *gin.Context) {
	scmt := comment.SubComment{}
	err := c.ShouldBind(&scmt)
	if err != nil {
		c.String(400, "Wrong param")
		return
	}
	cmtID := c.PostForm("commentID")
	if cmtID == "" {
		c.String(400, "Wrong comment id")
		return
	}
	scmt.CreateDate = time.Now().Format(time.RFC3339)

	err = comment.AddSubComment(cmtID, scmt)
	if err != nil {
		c.String(500, "Add sub comment failed")
		return
	}
	c.String(200, "Upload succeeded")
}

func md5String(original string) string {
	h := md5.New()
	io.WriteString(h, original)
	return hex.EncodeToString(h.Sum(nil))
}
