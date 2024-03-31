package mocks

import (
	"TaskManger/domain"
	"context"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) Create(c context.Context, user *domain.User) (*domain.User, error) {
    args := u.Called(c, user)
    return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepositoryMock) Delete(c context.Context, user *domain.User) (*domain.User, error) {
    args := u.Called(c, user)
    return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepositoryMock) GetById(c context.Context, id string) (*domain.User, error) {
    args := u.Called(c, id)
    return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepositoryMock) GetByUsername(c context.Context, username string) (*domain.User, error) {
    args := u.Called(c, username)
    return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepositoryMock) UpdatePassword(c context.Context, user *domain.User) (*domain.User, error) {
    args := u.Called(c, user)
    return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepositoryMock) UpdateUsername(c context.Context, user *domain.User) (*domain.User, error) {
    args := u.Called(c, user)
    return args.Get(0).(*domain.User), args.Error(1)
}

func NewUserRepositoryMock() domain.UserRepository {
	return &UserRepositoryMock{}
}



