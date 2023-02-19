package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func sendForm(user string, email string, pass string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("insert into users(username, email, password) values(?,?,?)", user, email, pass)
	if err != nil {
		panic(err)
	}
}
