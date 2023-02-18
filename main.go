package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type data struct {
	Username string
	Email    string
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./public/mainpage.html")
	db, _ := sql.Open("sqlite3", "/uses.db")
	db.Close()
	data := data{Username: "lily", Email: "hello"}
	tmp.Execute(w, data)
}
func serveSignup(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/signup.html")
}
func signUpFunc(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	email := r.FormValue("email")
	nonEncrypt := r.FormValue("pass")
	pass, _ := bcrypt.GenerateFromPassword([]byte(nonEncrypt), bcrypt.DefaultCost)
	sendForm(user, email, string(pass))
}

func loginFunc(w http.ResponseWriter, r *http.Request) {
	checkUsername(r.FormValue("user"))
	checkPassword(w, r, r.FormValue("pass"))
}
func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/login.html")
}
func main() {
	http.HandleFunc("/login/validation/", loginFunc)
	http.HandleFunc("/login/", serveLogin)
	http.HandleFunc("/signup/new/", signUpFunc)
	http.HandleFunc("/signup/", serveSignup)
	http.ListenAndServe(":4000", nil)
}
