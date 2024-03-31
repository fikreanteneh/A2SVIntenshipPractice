package test

import (
	"TaskManger/config"
	"TaskManger/domain"
	"TaskManger/models"
	"TaskManger/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUseCase_Login(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	env := &config.Environment{}

	userUC := usecase.NewUserUseCase(mockUserRepo, env, time.Second)

	payload := &models.UserCreate{
		Username: "testuser",
		Password: "testpassword",
	}

	mockUserRepo.On("GetByUsername", mock.Anything, payload.Username).Return(nil, errors.New("User Not Found"))

	token, err := userUC.Login(context.Background(), payload)

	assert.Error(t, err)
	assert.Equal(t, "", token)

	mockUserRepo.AssertExpectations(t)
}


func TestUserUseCase_Register(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	env := &config.Environment{} 

	userUC := usecase.NewUserUseCase(mockUserRepo, env, time.Second)

	payload := &models.UserCreate{
		Username: "testuser",
		Password: "testpassword",
	}

	mockUserRepo.On("GetByUsername", mock.Anything, payload.Username).Return(nil, nil)
	mockUserRepo.On("Create", mock.Anything, mock.Anything).Return(&domain.User{
		Username: payload.Username,
		Password: "hashed_password",
	}, nil)

	user, err := userUC.Register(context.Background(), payload)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, payload.Username, user.Username)

	mockUserRepo.AssertExpectations(t)
}
