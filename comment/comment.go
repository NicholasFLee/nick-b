package comment

import (
	"database/sql"
	"log"

	"github.com/nicholasflee/nick-b/db"
)

// Comment struct
type Comment struct {
	CommentID string `json:"commentID" form:"commentID"`
	// ArticleID   string `json:"articleID" form:"articleID" binding:"required"`
	Content     string `json:"content" form:"content" binding:"required"`
	CreateDate  string `json:"createDate" form:"createDate"`
	AuthorName  string `json:"authorName" from:"authorName" binding:"required"`
	IPAddress   string
	SubComments []SubComment `json:"subComments"`
}

// SubComment a property of Comment
type SubComment struct {
	AuthorName string `json:"authorName" form:"authorName" binding:"required"`
	CreateDate string `json:"createDate" form:"createDate"`
	Content    string `json:"content" form:"content" binding:"required"`
}

func init() {
	if err := createTables(); err != nil {
		log.Fatal(err)
	}
}

// GetComments by article id
func GetComments(atcID string) (cmts []Comment, err error) {
	selectCmts := `
		SELECT * 
		FROM comments
		WHERE ArticleID = ?
	`
	rows, err := db.DB.Query(selectCmts, atcID)
	if err != nil {
		return
	}
	cmts = []Comment{}
	for rows.Next() {
		var cmt Comment
		err = rows.Scan(nil, &cmt.CommentID, nil, &cmt.Content, &cmt.CreateDate, &cmt.AuthorName, &cmt.IPAddress)
		if err != nil {
			return
		}
		// get sub-comments of this comment
		var scmts []SubComment
		scmts, err = getSubComments(cmt.CommentID)
		if err != nil {
			return
		}
		cmt.SubComments = scmts
		cmts = append(cmts, cmt)
	}
	return
}

func getSubComments(cmtID string) (scmts []SubComment, err error) {
	selectScmts := `
		SELECT AuthorName, CreateDate, Content 
		From subcomments
		WHERE CommentID = ?
	`
	var rows *sql.Rows
	rows, err = db.DB.Query(selectScmts, cmtID)
	if err != nil {
		return
	}
	scmts = []SubComment{}
	for rows.Next() {
		scmt := SubComment{}
		err = rows.Scan(&scmt.AuthorName, &scmt.CreateDate, &scmt.Content)
		if err != nil {
			return
		}
		scmts = append(scmts, scmt)
	}
	return
}

// AddComment to a article
func AddComment(atcID string, cmt Comment) (err error) {
	insertCmt := `
		INSERT INTO comments
		(CommentID, ArticleID, Content, CreateDate, AuthorName, IPAddress)
		VALUES(?, ?, ?, ?, ?, ?)`
	_, err = db.DB.Exec(insertCmt, cmt.CommentID, atcID, cmt.Content, cmt.CreateDate, cmt.AuthorName, cmt.IPAddress)
	return
}

// AddSubComment add a sub-comment to a comment
func AddSubComment(cmtID string, scmt SubComment) (err error) {
	insertScmt := `
		INSERT INTO subcomments
		(CommentID, AuthorName, CreateDate, Content)
		VALUES(?, ?, ?, ?)
	`
	_, err = db.DB.Exec(insertScmt, cmtID, scmt.AuthorName, scmt.CreateDate, scmt.Content)
	return
}

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
			Content TEXT,
			PRIMARY KEY (ID)
		)
	`)
	return
}
