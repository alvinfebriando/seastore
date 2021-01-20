package domain

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("no user found")
)

type ErrorMessage struct {
	Message string
}

func (e ErrorMessage) Error() string {
	return e.Message
}
