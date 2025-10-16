package main

import (
	"PlayHouse/handlers"
	"PlayHouse/user"
	"net/http"
)

var usersList = make(map[string]*user.User)

func main() {
	http.HandleFunc("/user/new", handlers.HandlerNewUser(usersList))
	http.HandleFunc("/user/money", handlers.HandlerChangeBalance(usersList))
	http.ListenAndServe(":8080", nil)
}
