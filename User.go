package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type User struct {
	Username, Password, Email string
}

func NewUserFromJSON(req *http.Request) (user *User, err error) {
	decoder := json.NewDecoder(req.Body)
	user = new(User)
	err = decoder.Decode(user)
	return
}

func (user User) VerifyLogin() bool {
	return (user.VerifyUsername() || user.VerifyEmail()) && user.VerifyPassword()
}

func (user User) VerifyRegister() bool {
	return user.VerifyUsername() && user.VerifyEmail() && user.VerifyPassword()
}

func (user User) VerifyEmail() bool {
	match, err := regexp.MatchString("^.+@.+\\..+$", user.Email)
	return match && err == nil
}

func (user User) VerifyPassword() bool {
	return len(user.Password) >= 6
}

func (user User) VerifyUsername() bool {
	return user.Username != ""
}
