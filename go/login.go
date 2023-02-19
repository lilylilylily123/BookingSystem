package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func CheckUsername(username string) {
	db, _ := sql.Open("sqlite3", "./uses.db")
	user := db.QueryRow("select username from users where username= ?", username)
	temp := ""
	err := user.Scan(&temp)
	catch(err)
	if temp != "" {
		log.Println("Username is registered")
		return
	} else {
		log.Printf("Username %v is not registered.", username)
	}
	return
}

func CheckPassword(w http.ResponseWriter, r *http.Request, password string) {
	db, _ := sql.Open("sqlite3", "./uses.db")
	var hashed string
	err := db.QueryRow("select password from users where username=?",
		r.FormValue("user")).Scan(&hashed)
	if err != nil {
		log.Println("Password not registered")
	} else {
		encryptPass := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
		if encryptPass != nil {
			log.Println("not valid")
		} else {
			cookie := http.Cookie{Name: "username", Value: r.FormValue("user"), Path: "/"}
			http.SetCookie(w, &cookie)
			log.Println("Pass is registered")
		}
	}
	return
}
