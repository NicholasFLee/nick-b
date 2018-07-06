package comment

import (
	"fmt"
	"log"

	"github.com/nicholasflee/nick-b/db"
)

// Comment struct
type Comment struct {
	CommentID   string `json:"commentID" form:"commentID"`
	ArticleID   string `json:"articleID" form:"articleID" binding:"required"`
	Content     string `json:"content" form:"content" binding:"required"`
	CreateDate  string `json:"createDate" form:"createDate"`
	AuthorName  string `json:"authorName" from:"authorName" binding:"required"`
	IPAddress   string
	SubComments []SubComment `json:"subComments"`
}

// SubComment a property of Comment
type SubComment struct {
	CommentID  string `json:"commentID" form:"commentID"`
	AuthorName string `json:"authorName" form:"authorName"`
	CreateDate string `json:"createDate" form:"createDate"`
}

func init() {
	if err := createTables(); err != nil {
		log.Fatal(err)
	}
}

// GetComments by article id
func GetComments(atcID string) (cmts []Comment, err error) {
	selectCmts := fmt.Sprintf(`
		SELECT * FROM comments
		WHERE ArticleID='%s'
	`, atcID)
	rows, err := db.DB.Query(selectCmts)
	if err != nil {
		return
	}
	cmts = []Comment{}
	for rows.Next() {
		var cmt Comment
		err = rows.Scan(nil, &cmt.CommentID, &cmt.ArticleID, &cmt.Content, &cmt.CreateDate, &cmt.AuthorName, &cmt.IPAddress)
		if err != nil {
			return
		}
		cmts = append(cmts, cmt)
	}
	return
}

// AddComment to a article
func AddComment(cmt Comment) (err error) {
	insertCmt := `
		INSERT INTO comments
		(CommentID, ArticleID, Content, CreateDate, AuthorName, IPAddress)
		VALUES(?, ?, ?, ?, ?, ?)`
	_, err = db.DB.Exec(insertCmt, cmt.CommentID, cmt.ArticleID, cmt.Content, cmt.CreateDate, cmt.AuthorName, cmt.IPAddress)
	return
}

// AddSubComment add a sub-comment to a comment
func AddSubComment(scmt SubComment) (err error) {
	insertScmt := `
		INSERT INTO subcomments
		(CommentID, AuthorName, CreateDate)
		VALUES(?, ?, ?)
	`
	_, err = db.DB.Exec(insertScmt, scmt.CommentID, scmt.AuthorName, scmt.CreateDate)
	return
}

// func GetSubComment needed?

func createTables() (err error) {
	_, err = db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS comments(
			ID INT(11) NOT NULL AUTO_INCREMENT,
			CommentID varchar(100),
			ArticleID varchar(100),
			Content TEXT,
			CreateDate varchar(100),
			AuthorName varchar(100),
			IPAddress varchar(100),
			PRIMARY KEY (ID)
		)
	`)
	if err != nil {
		return
	}
	_, err = db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS subcomments(
			ID INT(11) NOT NULL AUTO_INCREMENT,
			CommentID varchar(100),
			AuthorName varchar(100),
			CreateDate varchar(100),
			PRIMARY KEY (ID)
		)
	`)
	return
}
