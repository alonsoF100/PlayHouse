package handlers

import (
	"PlayHouse/user"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerNewUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(user.NewUser); err != nil {
		fmt.Println("err ", err)
		return
	}
	fmt.Println(user)
}
