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

// TestUserUseCase_Login tests the Login method of UserUseCase.
func TestUserUseCase_Login(t *testing.T) {
	// Mock User Repository
	mockUserRepo := new(MockUserRepository)
	env := &config.Environment{} // You might want to initialize this with appropriate values

	// Initialize UserUseCase
	userUC := usecase.NewUserUseCase(mockUserRepo, env, time.Second)

	// Mock data
	payload := &models.UserCreate{
		Username: "testuser",
		Password: "testpassword",
	}

	// Mock behavior
	mockUserRepo.On("GetByUsername", mock.Anything, payload.Username).Return(nil, errors.New("User Not Found"))

	// Test the method
	token, err := userUC.Login(context.Background(), payload)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "", token) // Token should be empty if login fails

	mockUserRepo.AssertExpectations(t)
}


// TestUserUseCase_Register tests the Register method of UserUseCase.
func TestUserUseCase_Register(t *testing.T) {
	// Mock User Repository
	mockUserRepo := new(MockUserRepository)
	env := &config.Environment{} // You might want to initialize this with appropriate values

	// Initialize UserUseCase
	userUC := usecase.NewUserUseCase(mockUserRepo, env, time.Second)

	// Mock data
	payload := &models.UserCreate{
		Username: "testuser",
		Password: "testpassword",
	}

	// Mock behavior
	mockUserRepo.On("GetByUsername", mock.Anything, payload.Username).Return(nil, nil) // User does not exist
	mockUserRepo.On("Create", mock.Anything, mock.Anything).Return(&domain.User{
		Username: payload.Username,
		Password: "hashed_password",
	}, nil)

	// Test the method
	user, err := userUC.Register(context.Background(), payload)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, payload.Username, user.Username)

	mockUserRepo.AssertExpectations(t)
}
