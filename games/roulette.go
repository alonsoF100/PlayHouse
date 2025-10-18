package games

import (
	"PlayHouse/user"
	"fmt"
	"math/rand"
	"time"
)

func spin() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(38)
}

func Roulette(user *user.User, color string, bet int) {

	var winColor string

	if err := user.SetBalance(user.GetBalance() - bet); err != nil {
		fmt.Println("err =", err)
	}
	value := spin()

	switch {
	case value == 0:
		winColor = "green"
		fmt.Println(" green is color of the day! ")

	case value%2 == 0:
		winColor = "red"
		fmt.Println(" red is color of the day! ")
	default:
		winColor = "black"
		fmt.Println(" black is color of the day! ")
	}
	if color == winColor {
		switch color {
		case "red", "black":
			user.SetBalance(user.GetBalance() + bet*2)
			fmt.Println("ur lucky, now ur balace ", user.GetBalance())
			user.SetRating(user.GetRating() - 1)
		case "green":
			user.SetBalance(user.GetBalance() + bet*34)
			fmt.Println("ur lucky, now ur balace ", user.GetBalance())
			user.SetRating(user.GetRating() - 5)
		}

	} else {
		fmt.Println("better luck next time, now ur balace ", user.GetBalance())
		user.SetRating(user.GetRating() + 1)
	}
}
