package article

import (
	"fmt"
	"log"

	"github.com/nicholasflee/nick-b/db"
)

// Article struct will saved into 3 tables
type Article struct {
	ArticleID      string   `json:"aticleID" form:"id" binding:"required"`
	Title          string   `json:"title" form:"title" binding:"required"`
	CreateDate     string   `json:"createDate" form:"createDate"`
	Categories     []string `json:"categories" form:"categories" binding:"required"`
	Content        string   `json:"content" form:"content" binding:"required"`
	PreviewContent string   `json:"previewContent" form:"previewContent" binding:"required"`
}

type categories struct {
	// two primary key can ensure that `A C` group be unique.
	ArticleID    string
	CategoryName string
}

func init() {
	err := createTables()
	if err != nil {
		log.Fatal(err)
	}
}

// GetArticle by id
func GetArticle(id string) (err error) {
	// ...
	return nil
}

// GetArticlePreviews from and to
func GetArticlePreviews(page, perPage int) (atcs []Article, err error) {
	selectAtcs := fmt.Sprintf("SELECT * FROM articles OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", (page-1)*perPage, perPage)
	rows, err := db.DB.Query(selectAtcs)
	if err != nil {
		return
	}
	atcs = []Article{}
	if rows.Next() {
		var a Article
		err = rows.Scan(&a.ArticleID, &a.Title, &a.CreateDate, &a.Content, &a.PreviewContent)
		if err != nil {
			return
		}
		atcs = append(atcs, a)
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

	insertAtc := `INSERT INTO articles(ArticleID, Title, CreateDate, Content, PreviewContent) 
		  VALUES(?, ?, ?, ?, ?)`
	insertCts := `INSERT INTO categories(ArticleID, CategoryName) VALUES(?, ?)`

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
	_, err = db.DB.Exec(`CREATE TABLE IF NOT EXISTS articles 
		( ID INT(11) NOT NULL AUTO_INCREMENT, ArticleID varchar(100), Title varchar(100), 
		CreateDate varchar(100), Content TEXT, PreviewContent TEXT, PRIMARY KEY (ID) )`)
	if err != nil {
		return
	}
	_, err = db.DB.Exec(`CREATE TABLE IF NOT EXISTS categories
			( ArticleID varchar(100), CategoryName varchar(100), PRIMARY KEY (ArticleID, CategoryName) )`)
	return
}
