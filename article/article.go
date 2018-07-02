package article

import (
	"database/sql"
	"fmt"
	"log"

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
	PreviewContent string `json:"previewContent" form:"previewContent" binding:"required"`
}

func init() {
	err := createTables()
	if err != nil {
		log.Fatal(err)
	}
}

// GetArticle by id
func GetArticle(id string) (a Article, err error) {
	selectAtc := fmt.Sprintf(`
		SELECT ArticleID, Title, CreateDate, Content, PreviewContent 
		FROM myblog.articles 
		WHERE ArticleID='%s'`, id)
	err = db.DB.QueryRow(selectAtc).Scan(&a.ArticleID, &a.Title, &a.CreateDate, &a.Content, &a.PreviewContent)
	if err != nil {
		return
	}
	ctgs, err := gerCategoriesByArticleID(a.ArticleID)
	if err != nil {
		return
	}
	a.Categories = ctgs
	return
}

// GetArticlePreviews by page and perPage
func GetArticlePreviews(page, perPage int) (atcs []Article, err error) {
	selectAtcs := fmt.Sprintf(`
		SELECT ArticleID, Title, CreateDate, PreviewContent
		FROM myblog.articles 
		ORDER BY ID DESC
		LIMIT %d
		OFFSET %d`, perPage, (page-1)*perPage)

	rows, err := db.DB.Query(selectAtcs)
	if err != nil {
		return
	}
	atcs = []Article{}
	for rows.Next() {
		var a Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.CreateDate, &a.PreviewContent)
		if err != nil {
			return
		}
		// now add categories
		var ctgs []string
		ctgs, err = gerCategoriesByArticleID(a.ArticleID)
		if err != nil {
			return
		}
		a.Categories = ctgs
		atcs = append(atcs, a)
	}
	return
}

func gerCategoriesByArticleID(id string) (ctgs []string, err error) {
	selectCtg := fmt.Sprintf(`
		SELECT CategoryName
		FROM myblog.categories 
		WHERE ArticleID='%s'`, id)

	var ctgRows *sql.Rows
	ctgRows, err = db.DB.Query(selectCtg)
	if err != nil {
		return
	}
	ctgs = []string{}
	for ctgRows.Next() {
		var s string
		err = ctgRows.Scan(&s)
		if err != nil {
			return
		}
		ctgs = append(ctgs, s)
	}
	return
}

// InsertArticle insert into db
func InsertArticle(a Article) (err error) {
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
	tx.Commit()
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
