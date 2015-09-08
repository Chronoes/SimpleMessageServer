package main

import (
	// "crypto/md5"
	// "encoding/json"
	// "database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Register works")
}

func LoginUser(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Login works")
}

func main() {
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/login", LoginUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Errorf("Server error: %s", err)
	}
}
