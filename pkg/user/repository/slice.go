package repository

import (
	"github.com/alvinfebriando/seastore/pkg/domain"
	"github.com/google/uuid"
)

type repo struct {
	db []domain.User
}

func NewSliceRepository() *repo {
	return &repo{}
}

func (r *repo) Count() int {
	return len(r.db)
}

func (r *repo) Create(newUser domain.User) {
	r.db = append(r.db, newUser)
}

func (r *repo) FindByID(id uuid.UUID) (domain.User, error) {
	for _, v := range r.db {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}
func (r *repo) FindByEmail(email string) (domain.User, error) {
	for _, v := range r.db {
		if v.Email == email {
			return v, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}
