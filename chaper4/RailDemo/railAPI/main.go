package main

import (
	"database/sql"
	"log"
	"restful_with_golang/chaper4/RailDemo/dbutils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(db)
}
