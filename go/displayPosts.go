package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type allPost struct {
	Title  string
	Desc   string
	Id     int
	Author string
}

func displayAllPosts() {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	all := []allPost{}
	scan, err := db.Query("SELECT * from posts")
	catch(err)
	for scan.Next() {
		var p allPost
		err := scan.Scan(&p.Title, &p.Desc, &p.Id, &p.Author)
		catch(err)
		all = append(all, p)
	}
}
