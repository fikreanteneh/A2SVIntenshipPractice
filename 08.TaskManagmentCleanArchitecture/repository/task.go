package repository

import (
	"TaskManger/domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &TaskRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.TaskRepository.
func (t *TaskRepository) Create(c context.Context, task *domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// Delete implements domain.TaskRepository.
func (t *TaskRepository) Delete(c context.Context, taskId string) (domain.Task, error) {
	panic("unimplemented")
}

// GetById implements domain.TaskRepository.
func (t *TaskRepository) GetById(c context.Context, taskId string) (domain.Task, error) {
	panic("unimplemented")
}

// GetByUserId implements domain.TaskRepository.
func (t *TaskRepository) GetByUserId(c context.Context, userId string) ([]domain.Task, error) {
	panic("unimplemented")
}

// Update implements domain.TaskRepository.
func (t *TaskRepository) Update(c context.Context, task *domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

