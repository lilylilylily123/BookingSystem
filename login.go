package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func checkUsername(username string) {
	db, _ := sql.Open("sqlite3", "./uses.db")
	user := db.QueryRow("select username from users where username= ?", username)
	temp := ""
	user.Scan(&temp)
	if temp != "" {
		log.Println("Username is registered")
		return
	} else {
		log.Printf("Username %v is not registered.", username)
	}
}

func checkPassword(w http.ResponseWriter, r *http.Request, password string) {
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
			expires := time.Now().Add(time.Minute * 5)
			_ = fmt.Sprintf("Login expires in: %v minutes\n", expires)
			cookie := http.Cookie{Name: "loggedIn", Value: "true", Path: "/", Expires: expires}
			http.SetCookie(w, &cookie)
			log.Println("Pass is registered")
		}
	}
	return
}
