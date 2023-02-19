package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func create_post(post_title string, post_desc string, AID int, username string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	exec, err := db.Prepare("INSERT INTO posts(post_title, post_desc, author_id, username) values(?, ?, ?, ?)")
	catch(err)
	exec.Exec(post_title, post_desc, AID, username)
}
