package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB myblog db
var DB *sql.DB

func init() {
	db, err := OpenDB("myblog")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

// OpenDB a db
func OpenDB(name string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:MySqL123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		return
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		return
	}
	_, err = db.Exec("USE " + name)
	if err != nil {
		return
	}
	return
}
