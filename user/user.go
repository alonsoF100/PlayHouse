package user

import (
	"errors"
	"strings"
)

func (u user) NewUser(nickname string, password string, balance int, rating int) error {
	if strings.TrimSpace(nickname) == "" {
		return errors.New("nickname cannot be empty")
	}
	if len(password)<5{
		return errors.New("password so short")
	}
	if balance<0
}

type user struct {
	nickname string
	password string
	balance  int
	rating   int
}

