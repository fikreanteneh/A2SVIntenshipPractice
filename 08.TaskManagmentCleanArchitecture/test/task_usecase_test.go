package test

import (
    "TaskManger/config"
    "TaskManger/domain"
    "TaskManger/models"
    "TaskManger/usecase"
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "TaskManger/mocks"
    "testing"
    "time"
)

func TestTaskUseCase_Create(t *testing.T) {
    mockUserRepo := new(mocks.UserRepositoryMock)
    mockTaskRepo := new(mocks.TaskRepositoryMock)
    mockTask := &domain.Task{
        Title:       "Test Title",
        Description: "Test Description",
        Status:      false,
        DueDate:     time.Now(),
        UserId:      "TestUserId",
    }
    mockUser := &domain.User{
        Id:       "TestUserId",
        Username: "TestUsername",
        Password: "TestPassword",
    }
    mockUserRepo.On("GetByUsername", mock.Anything, "TestUsername").Return(mockUser, nil)
    mockTaskRepo.On("Create", mock.Anything, mock.Anything).Return(mockTask, nil)

    usecase := usecase.NewTaskUseCase(mockTaskRepo, mockUserRepo, &config.Environment{}, time.Second)

    task, err := usecase.Create(context.Background(), "TestUsername", &models.TaskCreate{
        Title:       "Test Title",
        Description: "Test Description",
        Status:      false,
        DueDate:     time.Now(),
    })

    assert.NoError(t, err)
    assert.NotNil(t, task)
    mockUserRepo.AssertExpectations(t)
    mockTaskRepo.AssertExpectations(t)
}

func TestTaskUseCase_Update(t *testing.T) {
    mockTaskRepo := new(mocks.TaskRepositoryMock)
    mockTask := &domain.Task{
        Title:       "Test Title",
        Description: "Test Description",
        Status:      false,
        DueDate:     time.Now(),
        UserId:      "TestUserId",
    }
    mockTaskRepo.On("Update", mock.Anything, mock.Anything).Return(mockTask, nil)

    usecase := usecase.NewTaskUseCase(mockTaskRepo, nil, &config.Environment{}, time.Second)
    task, err := usecase.Update(context.Background(), "TestUserId", &models.TaskUpdate{
        Title:       "Test Title",
        Description: "Test Description",
        Status:      false,
        DueDate:     time.Now(),
    })

    assert.NoError(t, err)
    assert.NotNil(t, task)
    mockTaskRepo.AssertExpectations(t)
}

func TestTaskUseCase_Delete(t *testing.T) {
    mockTaskRepo.On("Delete", mock.Anything, "TestUserId").Return(nil)

    err := usecase.Delete(context.Background(), "TestUserId")

    assert.NoError(t, err)
    mockTaskRepo.AssertExpectations(t)
}

func TestTaskUseCase_GetById(t *testing.T) {
    mockTask := &domain.Task{
        Title:       "Test Title",
        Description: "Test Description",
        Status:      false,
        DueDate:     time.Now(),
        UserId:      "TestUserId",
    }
    mockTaskRepo.On("GetById", mock.Anything, "TestUserId").Return(mockTask, nil)

    task, err := usecase.GetById(context.Background(), "TestUserId")

    assert.NoError(t, err)
    assert.NotNil(t, task)
    mockTaskRepo.AssertExpectations(t)
}

func TestTaskUseCase_GetByUserId(t *testing.T) {
    mockTasks := []*domain.Task{
        {
            Title:       "Test Title",
            Description: "Test Description",
            Status:      false,
            DueDate:     time.Now(),
            UserId:      "TestUserId",
        },
    }
    mockTaskRepo.On("GetByUserId", mock.Anything, "TestUserId").Return(mockTasks, nil)

    tasks, err := usecase.GetByUserId(context.Background(), "TestUserId")

    assert.NoError(t, err)
    assert.NotNil(t, tasks)
    mockTaskRepo.AssertExpectations(t)
}