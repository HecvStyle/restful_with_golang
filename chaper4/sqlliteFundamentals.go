// sqllite3 的基本数据操作

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}

	statement, err := db.Prepare("create table if not exists books(id integer primary key, isbn integer, author varchar(64), name varchar(64) NULL)")
	if err != nil {
		log.Println("error in careating table")
	} else {
		log.Println("Successfully created table books!")
	}
	statement.Exec()

}
