package article

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nicholasflee/nick-b/db"
)

// Article instance of a
type article struct {
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
	ArticleID  string
	CategoryID string
}

func init() {
	db, err := db.OpenDB("myblog")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	createArticleTable(db)
}

// GetArticleByID .
// func GetArticleByID(id string) Article {
// 	return article{Title: "title", Subtitle: "subtitle", CreateDate: "date", Categories: []string{""}}
// }

func createArticleTable(db *sql.DB) (err error) {
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS articles ( ID integer, Title varchar(32), Subitle varchar(32), CreateDate varchar(32), Categories )")
	return
}
