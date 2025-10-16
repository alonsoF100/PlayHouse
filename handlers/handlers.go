package handlers

import (
	"PlayHouse/user"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerNewUser(usersList map[string]*user.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// читаем json
		var dtoUser UserHandlerInfo
		if err := json.NewDecoder(r.Body).Decode(&dtoUser); err != nil {
			fmt.Println("err ", err)
			return
		}

		// Есть ли в мапе такой пользователь
		if _, exists := usersList[dtoUser.Nickname]; exists {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		// создание пользователя если нет
		validUser, err := user.NewUser(dtoUser.Nickname, dtoUser.Password, dtoUser.Balance, dtoUser.Rating)
		if err != nil {
			fmt.Println("err ", err)
			return
		}

		// добавляем в мапу
		usersList[validUser.GetNickname()] = validUser

		// статус код что все четко
		w.WriteHeader(http.StatusCreated)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "пользователь успешно добавлен")
		w.Header().Set("Content-Type", "application/json")
		// выввод в виде json добавленного пользователя
		responseUser := UserResponse{
			Nickname: validUser.GetNickname(),
			Balance:  validUser.GetBalance(),
			Rating:   validUser.GetRating(),
		}
		json.NewEncoder(w).Encode(responseUser)
	}
}

func HandlerChangeBalance(usersList map[string]*user.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userMoney UserHandlerMoney
		if err := json.NewDecoder(r.Body).Decode(&userMoney); err != nil {
			fmt.Println("err ", err)
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		user, exist := usersList[userMoney.Nickname]
		if !exist {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		if err := user.SetBalance(user.GetBalance() + userMoney.Money); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		responseUser := UserResponse{
			Nickname: user.GetNickname(),
			Balance:  user.GetBalance(),
			Rating:   user.GetRating(),
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "баланс успешно изменен")
		fmt.Println("баланс успешно изменен")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseUser)
	}
}
