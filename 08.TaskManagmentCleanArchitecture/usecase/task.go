package usecase

import (
	"TaskManger/domain"
	"TaskManger/models"
	"context"
	"time"
)

type TaskUseCase struct {
	TaskRepository domain.TaskRepository
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

// Create implements domain.TaskUseCase.
func (t *TaskUseCase) Create(c context.Context, username string, payload *models.TaskCreate) (*domain.Task, error) {
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	task := &domain.Task{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      payload.Status,
		DueDate:     payload.DueDate,
		UserId:      user.Id,
	}
	return t.TaskRepository.Create(c, task)
}

// Delete implements domain.TaskUseCase.
func (t *TaskUseCase) Delete(c context.Context, username string, taskId string) (*domain.Task, error) {
	//TODO: Authorization Handling
	_, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	task, err := t.TaskRepository.GetById(c, taskId)
	if err != nil {
		return nil, err
	}
	return t.TaskRepository.Delete(c, task)
	
}

// GetById implements domain.TaskUseCase.
func (t *TaskUseCase) GetById(c context.Context, username string, tasId string) (*domain.Task, error) {
	//TODO: Authorization Handling
	_, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	return t.TaskRepository.GetById(c, tasId)
}

// GetByUsername implements domain.TaskUseCase.
func (t *TaskUseCase) GetByUsername(c context.Context, username string) (*[]*domain.Task, error) {
	//TODO	: Authorization Handling
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	return t.TaskRepository.GetByUserId(c, user.Id)
}

// Update implements domain.TaskUseCase.
func (t *TaskUseCase) Update(c context.Context, username string, taskId string, payload *models.TaskUpdate) (*domain.Task, error) {
	//TODO: Authorization Handling
	_, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	task, err := t.TaskRepository.GetById(c, taskId)
	if err != nil {
		return nil, err
	}
	task.Title = payload.Title
	task.Description = payload.Description
	task.Status = payload.Status
	task.DueDate = payload.DueDate
	return t.TaskRepository.Update(c, task)
}

func NewTaskUseCase(tr domain.TaskRepository, ur domain.UserRepository, timeout time.Duration) domain.TaskUseCase {
	return &TaskUseCase{
		TaskRepository: tr,
		UserRepository: ur,
		contextTimeout: timeout,
	}
}
