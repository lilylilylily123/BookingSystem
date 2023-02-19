package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func getPost1() (string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc from posts order by random() limit 500")
	catch(err)
	var postTitle1, postDesc1 string
	for scan.Next() {
		scan.Scan(&postTitle1, &postDesc1)
	}
	//fmt.Println(postTitle1, postDesc1)
	return postTitle1, postDesc1
}
func getPost2() (string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc from posts order by random() limit 500")
	catch(err)
	var postTitle2, postdesc2 string
	for scan.Next() {
		scan.Scan(&postTitle2, &postdesc2)
	}
	//fmt.Println(postTitle2, postdesc2)
	return postTitle2, postdesc2
}
func getPost3() (string, string) {
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	scan, err := db.Query("SELECT post_title, post_desc from posts order by random() limit 500")
	catch(err)
	var postTitle3, postdesc3 string
	for scan.Next() {
		scan.Scan(&postTitle3, &postdesc3)
	}
	//fmt.Println(postTitle3, postdesc3)
	return postTitle3, postdesc3
}
