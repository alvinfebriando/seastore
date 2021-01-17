package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type UserRepository interface {
	Count() int
	Create(newUser User)
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
}

type UserService interface {
	Register() (User, error)
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
}

var (
	ErrUserNotFound = errors.New("no user found")
)
