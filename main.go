package main

import (
	// "crypto/sha1"
	// "database/sql"
	// "encoding/json"
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func RegisterUser(rw http.ResponseWriter, req *http.Request) {
	user, err := NewUserFromJSON(req)
	if err != nil {
		http.Error(rw, "Error decoding JSON "+err.Error(), 400)
		return
	} else if !user.VerifyRegister() {
		http.Error(rw, "Error: cannot verify registration", 401)
		return
	}
	log.Printf("Userdata recieved for registering: %s %s %s\n", user.Username, req.Method, req.URL.String())
	rw.Write([]byte("Register works"))
}

func LoginUser(rw http.ResponseWriter, req *http.Request) {
	user, err := NewUserFromJSON(req)
	if err != nil {
		http.Error(rw, "Error decoding JSON", 400)
		return
	} else if !user.VerifyLogin() {
		http.Error(rw, "Error: cannot verify login", 401)
		return
	}
	log.Printf("Userdata recieved for login: %s\n", user.Username)
	rw.Write([]byte("Login works"))
}

func main() {
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/login", LoginUser)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
