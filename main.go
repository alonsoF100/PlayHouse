package main

import (
	"PlayHouse/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/new", handlers.HandlerNewUser)

}
