package article

import (
	"database/sql"
	"log"

	"github.com/nicholasflee/nick-b/db"
)

// Article struct will saved into 3 tables
type Article struct {
	ID             string   `json:"id"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	CreateDate     string   `json:"createDate"`
	Categories     []string `json:"categories"`
	Content        string   `json:"content"`
	PreviewContent string   `json:"previewContent"`
}

type category struct {
	ID   string
	Name string
}

type categories struct {
	// two primary key can ensure that `A C` group be unique.
	ArticleID  string
	CategoryID string
}

func init() {
	err := createTables(db.DB)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertArticle insert into db
func InsertArticle(a Article) (err error) {
	return insertArticle(db.DB, a)
}

// db functions
func createTables(db *sql.DB) (err error) {
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS articles ( 
		ID varchar(100), Title varchar(100), Subtitle varchar(100), CreateDate varchar(100), Content TEXT, PreviewContent TEXT, PRIMARY KEY (ID) 
	)`)
	if err != nil {
		return
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS category 
		( ID INT(11) NOT NULL AUTO_INCREMENT, Name varchar(100), PRIMARY KEY (ID) )`)
	if err != nil {
		return
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories 
			( ArticleID varchar(100), CategoryID varchar(100), PRIMARY KEY (ArticleID, CategoryID) )`)
	return
}

func insertArticle(db *sql.DB, a Article) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	insertAtc := `INSERT INTO articles(ID, Title, Subtitle, CreateDate, Content, PreviewContent) 
		  VALUES(?, ?, ?, ?, ?, ?)`
	insertCtr := `INSERT INTO category(Name) VALUES(?)`
	insertCts := `INSERT INTO categories(ArticleID, CategoryID) VALUES(?, ?)`

	_, err = tx.Exec(insertAtc, a.ID, a.Title, a.Subtitle, a.CreateDate, a.Content, a.PreviewContent)
	if err != nil {
		return
	}
	for cate := range a.Categories {
		_, err = tx.Exec(insertCtr, cate)
		if err != nil {
			return
		}
		_, err = tx.Exec(insertCts, a.ID, cate)
		if err != nil {
			return
		}
	}
	tx.Commit()
	return
}
