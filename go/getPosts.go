package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func getPost1() (string, string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc, username from posts order by random() limit 500")
	catch(err)
	var postTitle1, postDesc1, author string
	for scan.Next() {
		scan.Scan(&postTitle1, &postDesc1, &author)
	}
	//fmt.Println(postTitle1, postDesc1)
	return postTitle1, postDesc1, author
}
func getPost2() (string, string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc, username from posts order by random() limit 500")
	catch(err)
	var postTitle2, postdesc2, author string
	for scan.Next() {
		scan.Scan(&postTitle2, &postdesc2, &author)
	}
	//fmt.Println(postTitle2, postdesc2)
	return postTitle2, postdesc2, author
}
func getPost3() (string, string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc, username from posts order by random() limit 500")
	catch(err)
	var postTitle3, postdesc3, author string
	for scan.Next() {
		scan.Scan(&postTitle3, &postdesc3, &author)
	}
	//fmt.Println(postTitle3, postdesc3)
	return postTitle3, postdesc3, author
}
