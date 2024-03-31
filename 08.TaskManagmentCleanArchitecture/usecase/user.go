package usecase

import (
	"TaskManger/config"
	"TaskManger/domain"
	"TaskManger/models"
	"TaskManger/utils"
	"context"
	"errors"
	"os/user"
	"time"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type UserUseCase struct {
	environment *config.Environment
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

// Login implements domain.UserUseCase.
func (u *UserUseCase) Login(c context.Context, payload *models.UserCreate) (string, error) {
	if payload.Username == "" || payload.Password == "" {
		return "", errors.New("Invalid Payload")
	}
	user, err := u.UserRepository.GetByUsername(c, payload.Username)
	if err != nil {
		return "", err
	}
	if user != nil {
		return "", errors.New("User Not Found")
	}
	match := utils.ComparePasswords(user.Password, payload.Password)
	if match != nil {
		return "", err
	}
	token, err :=  utils.TokenGenerate(user.Id, user.Username, u.environment.JwtSecret)
	return token, err


}

// Register implements domain.UserUseCase.
func (u *UserUseCase) Register(c context.Context, payload *models.UserCreate) (*domain.User, error) {
	if payload.Username == "" || payload.Password == "" {
		return nil, errors.New("Invalid Payload")
	}
	hashedPassword, err := utils.EncryptPassword(payload.Password)
	user	:= u.UserRepository.GetByUsername(c, payload.Username)
	if user != nil {
		return nil, errors.New("User Already Exists")
	}

	user, err := u.UserRepository.Create(c, &domain.User{
		Username: payload.Username,
		Password: hashedPassword,
	})
	return (*domain.User)(user), err
}

// Delete implements domain.UserUseCase.
func (u *UserUseCase) Delete(c context.Context, userId string) (*domain.User, error) {
	user, err := u.UserRepository.GetById(c, userId)
	user, err = u.UserRepository.Delete(c, user)
	return user, err
}

// GetById implements domain.UserUseCase.
func (u *UserUseCase) GetById(c context.Context, id string) (*domain.User, error) {
	user, err := u.UserRepository.GetById(c, id)
	return user, err
}

// GetByUsername implements domain.UserUseCase.
func (u *UserUseCase) GetByUsername(c context.Context, username string) (*domain.User, error) {
	user, err := u.UserRepository.GetByUsername(c, username)
	return user, err
}

// UpdatePassword implements domain.UserUseCase.
func (u *UserUseCase) UpdatePassword(c context.Context, userId string, payload *models.UserUpdatePassword) (*domain.User, error) {
	user, err := u.UserRepository.GetById(c, userId)
	hashedPassword, err := utils.EncryptPassword(payload.Password)
	user.Password = hashedPassword
	user, err = u.UserRepository.UpdatePassword(c, user)
	return user, err
}

// UpdateUsername implements domain.UserUseCase.
func (u *UserUseCase) UpdateUsername(c context.Context, userId string, payload *models.UserUpdateUsername) (*domain.User, error) {
	user, err := u.UserRepository.GetById(c, userId)
	user.Username = payload.Username
	user, err = u.UserRepository.UpdateUsername(c, user)
	return user, err
}

func NewUserUseCase(ur domain.UserRepository, env *config.Environment, timeout time.Duration) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: ur,
		environment: env,
		contextTimeout: timeout,
	}
}