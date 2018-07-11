package article

import (
	"database/sql"
	"log"

	"github.com/nicholasflee/nick-b/comment"
	"github.com/nicholasflee/nick-b/db"
)

// Article struct
type Article struct {
	ArticleID  string   `json:"articleID" form:"articleID"`
	Title      string   `json:"title" form:"title" binding:"required"`
	CreateDate string   `json:"createDate" form:"createDate"`
	Categories []string `json:"categories" form:"categories" binding:"required"`
	// in `article`
	Content string `json:"content" form:"content" binding:"required"`
	// in `preview articles`
	PreviewContent string            `json:"previewContent" form:"previewContent" binding:"required"`
	Comments       []comment.Comment `json:"comments"`
}

func init() {
	err := createTables()
	if err != nil {
		log.Fatal(err)
	}
}

// GetArticle by id
func GetArticle(id string) (a Article, err error) {
	selectAtc := `
		SELECT ArticleID, Title, CreateDate, Content, PreviewContent 
		FROM articles 
		WHERE ArticleID = ?`
	err = db.DB.QueryRow(selectAtc, id).Scan(&a.ArticleID, &a.Title, &a.CreateDate, &a.Content, &a.PreviewContent)
	if err != nil {
		return
	}
	var ctgs []string
	ctgs, err = gerCategoriesByArticleID(a.ArticleID)
	if err != nil {
		return
	}
	a.Categories = ctgs
	var cmts []comment.Comment
	cmts, err = comment.GetComments(a.ArticleID)
	if err != nil {
		return
	}
	a.Comments = cmts
	return
}

// GetArticlePreviews by page and perPage
func GetArticlePreviews(page, perPage int) (atcs []Article, err error) {
	selectAtcs := `
		SELECT ArticleID, Title, CreateDate, PreviewContent
		FROM articles 
		ORDER BY ID DESC
		LIMIT ?
		OFFSET ?
	`
	var rows *sql.Rows
	rows, err = db.DB.Query(selectAtcs, perPage, (page-1)*perPage)
	if err != nil {
		return
	}
	atcs = []Article{}
	defer rows.Close()
	for rows.Next() {
		var a Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.CreateDate, &a.PreviewContent)
		if err != nil {
			return
		}
		atcs = append(atcs, a)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	// now add categories
	for _, a := range atcs {
		var ctgs []string
		ctgs, err = gerCategoriesByArticleID(a.ArticleID)
		if err != nil {
			return
		}
		a.Categories = ctgs
	}
	return
}

func gerCategoriesByArticleID(id string) (ctgs []string, err error) {
	_, err = db.DB.Exec("USE myblog")
	if err != nil {
		return
	}
	selectCtg := `
		SELECT CategoryName
		FROM categories 
		WHERE ArticleID = ?`

	var ctgRows *sql.Rows
	ctgRows, err = db.DB.Query(selectCtg, id)
	if err != nil {
		return
	}
	ctgs = []string{}
	defer ctgRows.Close()
	for ctgRows.Next() {
		var s string
		err = ctgRows.Scan(&s)
		if err != nil {
			return
		}
		ctgs = append(ctgs, s)
	}
	err = ctgRows.Err()
	return
}

// AddArticle insert into db
func AddArticle(a Article) (err error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	insertAtc := `
		INSERT INTO articles
		(ArticleID, Title, CreateDate, Content, PreviewContent) 
		VALUES(?, ?, ?, ?, ?)`
	insertCts := `
		INSERT INTO categories(ArticleID, CategoryName) 
		VALUES(?, ?)`

	_, err = tx.Exec(insertAtc, a.ArticleID, a.Title, a.CreateDate, a.Content, a.PreviewContent)
	if err != nil {
		return
	}
	for _, cate := range a.Categories {
		_, err = tx.Exec(insertCts, a.ArticleID, cate)
		if err != nil {
			return
		}
	}
	err = tx.Commit()
	return
}

func createTables() (err error) {
	_, err = db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS articles(
			ID INT(11) NOT NULL AUTO_INCREMENT, 
			ArticleID varchar(100), 
			Title varchar(100), 
			CreateDate varchar(100), 
			Content TEXT, 
			PreviewContent TEXT, 
			PRIMARY KEY (ID)
		)
	`)
	if err != nil {
		return
	}
	_, err = db.DB.Exec(`
		CREATE TABLE IF NOT EXISTS categories( 
			ArticleID varchar(100), 
			CategoryName varchar(100), 
			PRIMARY KEY (ArticleID, CategoryName) 
		)
	`)
	return
}
