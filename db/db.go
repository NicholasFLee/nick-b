package db

import (
	"database/sql"
)

// OpenDB a db
func OpenDB(name string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "nick:MySqL123456@tcp(127.0.0.1:3306)/")
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

// func (*DB) createTable(name string) error {
// 	_, err := db.Exec("CREATE TABLE IF NOT EXISTS ? ( id integer, name varchar(32) )", name)
// 	return err
// }
