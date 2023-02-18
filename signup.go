package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func sendForm(user string, email string, pass string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("insert into users(username, email, password) values(?,?,?)")
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(user, email, pass)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	log.Println("kaldadlald")
}
