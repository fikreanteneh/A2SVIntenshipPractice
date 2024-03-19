package usecase

import (
	"TaskManger/domain"
	"time"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}


func NewUserUseCase(ur domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: ur,
		contextTimeout: timeout,
	}
}

// GetById implements domain.UserUseCase.
func (u *UserUseCase) GetById(id string) (domain.User, error) {
	panic("unimplemented")
}

// Update implements domain.UserUseCase.
func (u *UserUseCase) Update(user *domain.User) (domain.User, error) {
	panic("unimplemented")
}



// Create implements domain.UserUseCase.
func (u *UserUseCase) Create(user *domain.User) (domain.User, error) {
	panic("unimplemented")
}

// GetByUsername implements domain.UserUseCase.
func (u *UserUseCase) GetByUsername(username string) (domain.User, error) {
	panic("unimplemented")
}
