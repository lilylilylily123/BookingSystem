package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

type Post struct {
	Username string
	Email    string
	Title1   string
	Desc1    string
	Title2   string
	Desc2    string
	Title3   string
	Desc3    string
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./public/mainpage.html")
	catch(err)
	db, _ := sql.Open("sqlite3", "./uses.db")
	cookie, _ := r.Cookie("username")
	value := cookie.Value
	user, err := db.Query("SELECT username, email from users where username=?", value)
	catch(err)
	var displayName, displayEmail string
	for user.Next() {
		err := user.Scan(&displayName, &displayEmail)
		catch(err)
	}
	content := Post{Username: displayName, Email: displayEmail}
	t1, d1 := getPost1()
	t2, d2 := getPost2()
	t3, d3 := getPost3()
	content = Post{Title1: t1, Desc1: d1, Title2: t2, Desc2: d2, Title3: t3, Desc3: d3}
	err = tmp.Execute(w, content)
	catch(err)

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
	http.Redirect(w, r, "/login/", 302)
}

func loginFunc(w http.ResponseWriter, r *http.Request) {
	CheckUsername(r.FormValue("user"))
	CheckPassword(w, r, r.FormValue("pass"))
}
func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/login.html")
}

func sendPost(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	var id int
	var username string
	scan, _ := db.Query("SELECT id, username from users where username=?", cookie.Value)
	for scan.Next() {
		err := scan.Scan(&id, &username)
		catch(err)
	}
	fmt.Println(id)
	create_post(r.FormValue("list-title"), r.FormValue("list-desc"), id, username)
}
func serveCreatePost(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/createpost.html")
}
func servePosts(w http.ResponseWriter, r *http.Request) {
	displayPosts(w)
}
func main() {
	http.HandleFunc("/book/show/", servePosts)
	http.HandleFunc("/book/new/", serveCreatePost)
	http.HandleFunc("/book/new/create/", sendPost)
	http.HandleFunc("/login/validation/", loginFunc)
	http.HandleFunc("/login/", serveLogin)
	http.HandleFunc("/signup/new/", signUpFunc)
	http.HandleFunc("/mainpage/", mainpage)
	http.HandleFunc("/signup/", serveSignup)
	err := http.ListenAndServe(":4000", nil)
	catch(err)
}
func catch(err error) {
	if err != nil {
		log.Println(err)
	}
}
