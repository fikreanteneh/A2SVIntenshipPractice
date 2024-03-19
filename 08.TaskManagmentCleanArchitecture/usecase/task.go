package usecase

import (
	"TaskManger/domain"
	"context"
	"time"
)

type TaskUseCase struct {
	TaskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUseCase(tr domain.TaskRepository, timeout time.Duration) domain.TaskUseCase {
	return &TaskUseCase{
		TaskRepository: tr,
		contextTimeout: timeout,
	}
}

// Create implements domain.TaskUseCase.
func (t *TaskUseCase) Create(c context.Context, task *domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// Delete implements domain.TaskUseCase.
func (t *TaskUseCase) Delete(c context.Context, taskId string) error {
	panic("unimplemented")
}

// GetById implements domain.TaskUseCase.
func (t *TaskUseCase) GetById(c context.Context, taskId string) (domain.Task, error) {
	panic("unimplemented")
}

// GetByUserId implements domain.TaskUseCase.
func (t *TaskUseCase) GetByUserId(c context.Context, userId string) ([]domain.Task, error) {
	panic("unimplemented")
}

// Update implements domain.TaskUseCase.
func (t *TaskUseCase) Update(c context.Context, task *domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

