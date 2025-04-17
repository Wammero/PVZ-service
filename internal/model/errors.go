package model

import (
	"fmt"
)

type Error struct {
	Message string `json:"message"`
}

type UserAlreadyExistsError struct {
	Email string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("пользователь с email %s уже зарегистрирован", e.Email)
}
