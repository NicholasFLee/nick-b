package db

import (
	"database/sql"
	"fmt"
	"log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB myblog db
var DB *sql.DB

func init() {
	d, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	err = d.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = d
	fmt.Println("opened db")
}

// OpenDB db
func openDB() (d *sql.DB, err error) {
	d, err = sql.Open("mysql", "root:MySqL123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		return
	}
	_, err = d.Exec("CREATE DATABASE IF NOT EXISTS myblog")
	if err != nil {
		return
	}
	_, err = d.Exec("USE myblog")
	fmt.Println(d)
	return
}
