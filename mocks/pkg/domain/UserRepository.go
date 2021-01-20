// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/alvinfebriando/seastore/pkg/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
	Cache map[uuid.UUID]domain.User
}

// Count provides a mock function with given fields:
func (_m *UserRepository) Count() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Create provides a mock function with given fields: newUser
func (_m *UserRepository) Create(newUser domain.User) {
	_m.Called(newUser)
	_m.Cache[newUser.ID] = newUser
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserRepository) FindByEmail(email string) (domain.User, error) {
	ret := _m.Called(email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *UserRepository) FindByID(id uuid.UUID) (domain.User, error) {
	ret := _m.Called(id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(uuid.UUID) domain.User); ok {
		r0 = rf(id)
	} else {
		// r0 = ret.Get(0).(domain.User)
		r0 = _m.Cache[id]
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindByUsername(username string) (domain.User, error) {
	ret := _m.Called(username)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
