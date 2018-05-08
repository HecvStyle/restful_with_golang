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

	//query
	rows, _ := db.Query("select id, name, author from books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d ,Book:%s,Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// update
	statement, _ = db.Prepare("update books set name = ? where id = ?")
	statement.Exec("The tale of two cities", 1)
	log.Println("successfully update the book in database!")

	// delete
	statement, _ = db.Prepare("delete from books where id =?")
	statement.Exec(1)
	log.Println("successfully delete the book in database!")
}
