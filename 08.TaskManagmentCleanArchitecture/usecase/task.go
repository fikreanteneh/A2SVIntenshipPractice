package usecase

import (
	"TaskManger/config"
	"TaskManger/domain"
	"TaskManger/models"
	"context"
	"errors"
	"time"
)

type TaskUseCase struct {
	environment *config.Environment
	TaskRepository domain.TaskRepository
	UserRepository domain.UserRepository
	contextTimeout time.Duration
}

func (t *TaskUseCase) Create(c context.Context, username string, payload *models.TaskCreate) (*domain.Task, error) {
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil || user == nil{
		return nil, err
	}
	if payload.Status == nil || payload.Title == "" || payload.DueDate == nil || payload.Description == "" {
		return nil, errors.New("Invalid Payload")
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
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	task, err := t.TaskRepository.GetById(c, taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != user.Id {
		return nil, errors.New("Unauthorized")
	}
	return t.TaskRepository.Delete(c, task)
	
}

// GetById implements domain.TaskUseCase.
func (t *TaskUseCase) GetById(c context.Context, username string, tasId string) (*domain.Task, error) {
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	task, err := t.TaskRepository.GetById(c, tasId)
	if task.UserId != user.Id {
		return nil, errors.New("Unauthorized")
	}
	return task, nil
}

// GetByUsername implements domain.TaskUseCase.
func (t *TaskUseCase) GetByUsername(c context.Context, username string) (*[]*domain.Task, error) {
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	return t.TaskRepository.GetByUserId(c, user.Id)
}

// Update implements domain.TaskUseCase.
func (t *TaskUseCase) Update(c context.Context, username string, taskId string, payload *models.TaskUpdate) (*domain.Task, error) {
	user, err := t.UserRepository.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	task, err := t.TaskRepository.GetById(c, taskId)
	if err != nil {
		return nil, err
	}
	if task.UserId != user.Id {
		return nil, errors.New("Unauthorized")
	}
	task.Title = payload.Title
	task.Description = payload.Description
	task.Status = payload.Status
	task.DueDate = payload.DueDate
	return t.TaskRepository.Update(c, task)
}

func NewTaskUseCase(tr domain.TaskRepository, ur domain.UserRepository, env *config.Environment, timeout time.Duration) domain.TaskUseCase {
	return &TaskUseCase{
		TaskRepository: tr,
		UserRepository: ur,
		environment: env,
		contextTimeout: timeout,
	}
}
