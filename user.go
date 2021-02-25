package player

import "fmt"

func NewErrorUserNotFound(id int) error {
	return fmt.Errorf("User with id #%d not found", id)
}

func NewErrorTokenNotFound(id uint64) error {
	return fmt.Errorf("Token with id #%d is invalid", id)
}

func NewErrorUserHasBeenCreated(id uint64) error {
	return fmt.Errorf("User with id #%d has been created", id)
}

type User struct {
	Id      uint64 `json:"id"`
	Balance int    `json:"balance"`
	Token   string `json:"token"`
}
