package main

import (
	// "crypto/sha1"
	"encoding/json"
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type User struct {
	Username, Password, Email string
}

func NewUserFromJSON(req *http.Request) (user *User, err error) {
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(user)
	return user, err
}

func RegisterUser(rw http.ResponseWriter, req *http.Request) {
	user, err := NewUserFromJSON(req)
	if err != nil {
		http.Error(rw, "Error decoding JSON", 400)
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
