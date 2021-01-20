package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Username string
}

type UserRepository interface {
	Count() int
	Create(newUser User)
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
}

type UserService interface {
	Register() (User, error)
	FindByID(id uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
}
