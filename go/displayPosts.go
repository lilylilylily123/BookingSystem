package main

import (
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

//type Post struct {
//	Username string
//	Email    string
//	Title1   string
//	Desc1    string
//	Title2   string
//	Desc2    string
//	Title3   string
//	Desc3    string
//}

func displayPosts(w http.ResponseWriter) {
	t1, d1 := getPost1()
	t2, d2 := getPost2()
	t3, d3 := getPost3()
	p := Post{Title1: t1, Desc1: d1, Title2: t2, Desc2: d2, Title3: t3, Desc3: d3}
	tmp, err := template.ParseFiles("./public/mainpage.html")
	catch(err)
	err = tmp.Execute(w, p)
	catch(err)
}
