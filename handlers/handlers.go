package handlers

import (
	"PlayHouse/user"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerNewUser(w http.ResponseWriter, r *http.Request) {
	var dtoUser UserHandlerInfo
	if err := json.NewDecoder(r.Body).Decode(&dtoUser); err != nil {
		fmt.Println("err ", err)
		return
	}
	user1, err := user.NewUser(dtoUser.nickname, dtoUser.password, dtoUser.balance, dtoUser.rating)
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	fmt.Println(user1)
}
