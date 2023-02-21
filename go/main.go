package main

import (
	"database/sql"
	"errors"
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
	Auth1    string
	Title2   string
	Desc2    string
	Auth2    string
	Title3   string
	Desc3    string
	Auth3    string
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./public/mainpage.html")
	catch(err)
	db, _ := sql.Open("sqlite3", "./uses.db")
	cookie, err := r.Cookie("username")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.ServeFile(w, r, "./public/uhoh.html")
			//http.Redirect(w, r, "/uhoh/", 302)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	} else if err == nil {
		value := cookie.Value
		user, err := db.Query("SELECT username, email from users where username=?", value)
		catch(err)
		var displayName, displayEmail string
		for user.Next() {
			err := user.Scan(&displayName, &displayEmail)
			catch(err)
		}
		t1, d1, a1 := getPost1()
		t2, d2, a2 := getPost2()
		t3, d3, a3 := getPost3()
		content := Post{Username: displayName, Email: displayEmail, Auth1: a1, Title1: t1, Desc1: d1, Auth2: a2, Title2: t2, Desc2: d2, Auth3: a3, Title3: t3, Desc3: d3}
		err = tmp.Execute(w, content)
		catch(err)
	}
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
	CheckUsername(w, r, r.FormValue("user"))
	CheckPassword(w, r, r.FormValue("pass"))
	http.Redirect(w, r, "/mainpage/", 302)
}
func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/login.html")
}

func sendPost(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.ServeFile(w, r, "./public/uhoh.html")
			//http.Redirect(w, r, "/uhoh/", 302)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	} else if err == nil {
		db, err := sql.Open("sqlite3", "./uses.db")
		catch(err)
		var id int
		var username string
		scan, _ := db.Query("SELECT id, username from users where username=?", cookie.Value)
		for scan.Next() {
			err := scan.Scan(&id, &username)
			catch(err)
		}
		//fmt.Println(id)
		create_post(r.FormValue("list-title"), r.FormValue("list-desc"), id, username)
		http.Redirect(w, r, "/mainpage/", 302)
	}
}
func serveCreatePost(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/createpost.html")
}
func servePosts(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./public/allPosts.html")
	catch(err)
	db, err := sql.Open("sqlite3", "./uses.db")
	catch(err)
	all := []allPost{}
	scan, err := db.Query("SELECT * from posts")
	catch(err)
	var title, desc, auth string
	var id int
	for scan.Next() {
		err := scan.Scan(&title, &desc, &id, &auth)
		catch(err)
		all = append(all, allPost{
			Title: title, Desc: desc, Id: id, Author: auth,
		})
	}
	err = tmp.Execute(w, all)
	catch(err)
}
func serveAcnt(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.ServeFile(w, r, "./public/uhoh.html")
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	} else if err == nil {
		db, err := sql.Open("sqlite3", "./uses.db")
		catch(err)
		scan, err := db.Query("SELECT username, email from users where username=?", c.Value)
		var u, e string
		for scan.Next() {
			err := scan.Scan(&u, &e)
			catch(err)
		}
		displayAcnt(w, u, e)
	}
}
func main() {
	http.HandleFunc("/account/settings/", serveAcnt)
	http.HandleFunc("/display/", servePosts)
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
