package service_test

import (
	"testing"

	mocks "github.com/alvinfebriando/seastore/mocks/pkg/domain"
	"github.com/alvinfebriando/seastore/pkg/domain"
	"github.com/alvinfebriando/seastore/pkg/user/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	t.Run("When Register called, it should return createdUser", func(t *testing.T) {
		repo := &mocks.UserRepository{Cache: make(map[uuid.UUID]domain.User)}
		s := service.NewService(repo)
		user := domain.User{Name: "TestName", Email: "TestEmail"}

		repo.On("Create", mock.Anything)
		repo.On("FindByID", mock.Anything).Return(user, nil)

		createdUser, err := s.Register(user)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, createdUser.ID)
	})
}

func TestFind(t *testing.T) {
	t.Run("When FindByID with wrong id, it should error", func(t *testing.T) {
		repo := &mocks.UserRepository{Cache: make(map[uuid.UUID]domain.User)}
		s := service.NewService(repo)
		user := domain.User{Name: "TestName", Email: "TestEmail"}

		repo.On("Create", mock.Anything)
		repo.On("FindByID", mock.Anything).Return(domain.User{}, domain.ErrUserNotFound)

		s.Register(user)
		result, err := s.FindByID(uuid.Nil)

		assert.Error(t, err)
		assert.Errorf(t, err, "no user found")
		assert.Equal(t, "", result.Name)
	})

	t.Run("When FindByEmail with valid email, it should return the user", func(t *testing.T) {
		repo := &mocks.UserRepository{Cache: make(map[uuid.UUID]domain.User)}
		s := service.NewService(repo)
		user := domain.User{Name: "TestName", Email: "TestEmail"}

		repo.On("Create", mock.Anything)
		repo.On("FindByID", mock.Anything).Return(user, nil)
		repo.On("FindByEmail", mock.Anything).Return(user, nil)

		s.Register(user)
		searchedUser, err := s.FindByEmail(user.Email)

		assert.NoError(t, err)
		assert.Equal(t, user.Email, searchedUser.Email)
	})

	t.Run("When FindByEmail with wrong email, it should error", func(t *testing.T) {
		repo := &mocks.UserRepository{}
		s := service.NewService(repo)

		repo.On("FindByEmail", mock.Anything).Return(domain.User{}, domain.ErrUserNotFound)

		result, err := s.FindByEmail("")

		assert.Error(t, err)
		assert.Errorf(t, err, "no user found")
		assert.Equal(t, "", result.Email)
	})
}
