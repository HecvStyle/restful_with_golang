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

	statement, _ = db.Prepare("insert into books (name, author, isbn) values(?,?,?)")
	statement.Exec("A table of two cities", "Charles Dikends", 140430547)
	log.Println("Inserted the book into database!")

	rows, _ := db.Query("select id, name, author frome books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d ,Bool:%s,Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}
}
