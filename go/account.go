package main

import (
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

type acntInfo struct {
	Email string
	User  string
}

func displayAcnt(w http.ResponseWriter, user string, email string) {
	tmp, err := template.ParseFiles("./public/account.html")
	catch(err)
	d := acntInfo{User: user, Email: email}
	err = tmp.Execute(w, d)
	catch(err)
}
