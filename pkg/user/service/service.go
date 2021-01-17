package service

import (
	"github.com/alvinfebriando/seastore/pkg/domain"
	"github.com/google/uuid"
)

type Service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(newUser domain.User) (domain.User, error) {
	newUser.ID = uuid.New()
	s.repo.Create(newUser)
	createdUser, _ := s.FindByID(newUser.ID)
	return createdUser, nil
}

func (s *Service) FindByID(id uuid.UUID) (domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *Service) FindByEmail(email string) (domain.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
