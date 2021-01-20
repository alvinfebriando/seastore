package repository_test

import (
	"testing"

	"github.com/alvinfebriando/seastore/pkg/domain"
	"github.com/alvinfebriando/seastore/pkg/user/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewRepo(t *testing.T) {
	repo := repository.NewSliceRepository()
	t.Run("Initialized repo should not be nil", func(t *testing.T) {
		assert.NotNil(t, repo)
	})

	t.Run("Size of empty db should return 0", func(t *testing.T) {
		c := repo.Count()
		assert.Equal(t, c, 0)
	})
}

func TestCreateAndFind(t *testing.T) {
	t.Run("FindByID should return a user with same id", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		id := uuid.New()
		newUser := domain.User{ID: id, Name: "TestName", Email: "TestEmail"}
		repo.Create(newUser)
		u, err := repo.FindByID(id)

		assert.NoError(t, err)
		assert.Equal(t, newUser.ID, u.ID)
		assert.Equal(t, 1, repo.Count())
	})

	t.Run("FindByID should return error if no user with that id ", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		u, err := repo.FindByID(uuid.New())

		assert.Error(t, err)
		assert.Errorf(t, err, "no user found")
		assert.Equal(t, uuid.Nil, u.ID)
	})

	t.Run("FindByEmail should return a user with same email", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		id := uuid.New()
		newUser := domain.User{ID: id, Name: "TestName", Email: "TestEmail"}
		repo.Create(newUser)
		u, err := repo.FindByEmail("TestEmail")

		assert.NoError(t, err)
		assert.Equal(t, newUser.ID, u.ID)
		assert.Equal(t, 1, repo.Count())
	})

	t.Run("FindByEmail should return error if no user with that email ", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		u, err := repo.FindByEmail("fakeemail")

		assert.Error(t, err)
		assert.Errorf(t, err, "no user found")
		assert.Equal(t, "", u.Email)
	})

	t.Run("FindByUsername should return a user with same username", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		id := uuid.New()
		newUser := domain.User{ID: id, Name: "TestName", Email: "TestEmail", Username: "TestUsername"}
		repo.Create(newUser)
		u, err := repo.FindByUsername("TestUsername")

		assert.NoError(t, err)
		assert.Equal(t, newUser.Username, u.Username)
		assert.Equal(t, 1, repo.Count())
	})

	t.Run("FindByUsername should return error if no user with that username ", func(t *testing.T) {
		repo := repository.NewSliceRepository()
		u, err := repo.FindByUsername("fakeUsername")

		assert.Error(t, err)
		assert.Errorf(t, err, "no user found")
		assert.Equal(t, "", u.Username)
	})
}

// func TestUpdate(t *testing.T) {
// 	repo := slice.NewSliceRepository()

// 	t.Run("Update in empty db should return error", func(t *testing.T) {
// 		repo.Update("1", &domain.User{})
// 	})

// 	t.Run("Update should update user with same id", func(t *testing.T) {
// 		newUser := &domain.User{"1", "A"}
// 		repo.Create(newUser)

// 		u, _ := repo.GetByID("1")
// 		assert.Equal(t, u.Name, "A")

// 		updatedUser := &domain.User{"1", "B"}
// 		err := repo.Update("1", updatedUser)

// 		assert.NoError(t, err)

// 		u, _ = repo.GetByID("1")
// 		assert.Equal(t, u.Name, "B")
// 	})

// 	t.Run("Update with wrong id should return an error", func(t *testing.T) {
// 		err := repo.Update("10", &domain.User{})
// 		assert.Error(t, err)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	repo := slice.NewSliceRepository()

// 	t.Run("Delete a user in empty db should return error", func(t *testing.T) {
// 		err := repo.Delete("1")
// 		assert.Error(t, err)
// 	})

// 	newUser := &domain.User{ID: "1", Name: "A"}
// 	repo.Create(newUser)

// 	t.Run("Delete a user with wrong id should return error", func(t *testing.T) {
// 		err := repo.Delete("10")
// 		assert.Error(t, err)
// 	})

// 	t.Run("Delete a user should delete user in db with same id", func(t *testing.T) {
// 		err := repo.Delete("1")
// 		assert.NoError(t, err)

// 		user, err := repo.GetByID("1")
// 		assert.Error(t, err)
// 		assert.Nil(t, user)
// 	})
// }
