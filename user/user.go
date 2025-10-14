package user

import (
	"errors"
	"strings"
)

type User struct {
	nickname string
	password string
	balance  int
	rating   int
}

// конструктор для пользователя
func NewUser(nickname string, password string, balance int, rating int) (*User, error) {
	// валидация имени
	validNickname := strings.TrimSpace(nickname)
	if validNickname == "" {
		return nil, errors.New("nickname cannot be empty")
	}

	if len([]rune(validNickname)) < 5 || len([]rune(validNickname)) > 10 {
		return nil, errors.New("nickname must be between 5 and 10 characters long")
	}

	// валидация пароля
	validPassword := strings.TrimSpace(password)
	validPasswordRune := []rune(validPassword)
	hasUpper := false
	count := 0
	for _, v := range validPasswordRune {
		if v >= 'A' && v <= 'Z' {
			hasUpper = true
		}

		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			count++
		}
	}

	if !hasUpper {
		return nil, errors.New("the password must contain a upper letter")
	}

	if count < 2 {
		return nil, errors.New("the password must contain at least 2 letters")
	}

	if len(validPasswordRune) < 10 {
		return nil, errors.New("the password must be at least 10 characters long")
	}

	// валидация баланса
	if balance < 0 {
		return nil, errors.New("the balance cannot be negative")
	}

	// валидация рейтинга
	if rating < 0 || rating > 10 {
		return nil, errors.New("the rating must be from 0 to 10")
	}

	u := &User{
		nickname: validNickname,
		password: validPassword,
		balance:  balance,
		rating:   rating,
	}

	return u, nil
}

// геттеры для пользователя
func (u *User) GetNickname() string {
	return u.nickname
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetBalance() int {
	return u.balance
}

func (u *User) GetRating() int {
	return u.rating
}

// сеттеры для пользователя
func (u *User) SetNickname(newNickname string) error {
	validNickname := strings.TrimSpace(newNickname)
	if validNickname == "" {
		return errors.New("nickname cannot be empty")
	}

	if len([]rune(validNickname)) < 5 || len([]rune(validNickname)) > 10 {
		return errors.New("nickname must be between 5 and 10 characters long")
	}

	u.nickname = validNickname
	return nil
}

func (u *User) SetPassword(newPassword string) error {
	validPassword := strings.TrimSpace(newPassword)
	validPasswordRune := []rune(validPassword)
	hasUpper := false
	count := 0
	for _, v := range validPasswordRune {
		if v >= 'A' && v <= 'Z' {
			hasUpper = true
		}

		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			count++
		}
	}

	if !hasUpper {
		return errors.New("the password must contain a upper letter")
	}

	if count < 2 {
		return errors.New("the password must contain at least 2 letters")
	}

	if len(validPasswordRune) < 10 {
		return errors.New("the password must be at least 10 characters long")
	}

	u.password = validPassword
	return nil
}

func (u *User) SetBalance(newBalance int) error {
	if newBalance < 0 {
		return errors.New("the balance cannot be negative")
	}

	u.balance = newBalance
	return nil
}

func (u *User) SetRating(newRating int) error {
	if newRating < 0 || newRating > 10 {
		return errors.New("the rating must be from 0 to 10")
	}

	u.rating = newRating
	return nil
}
