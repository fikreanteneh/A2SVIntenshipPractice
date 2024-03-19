package domain

import (
	"context"
	"time"
)


type Task struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      bool   `json:"status" bson:"status"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
}


type TaskRepository interface {
	Create(c context.Context, task *Task) (Task, error)
	Update(c context.Context, task *Task) (Task, error)
	Delete(c context.Context, taskId string) (Task, error)
	GetByUserId(c context.Context, userId string) ([]Task, error)
	GetById(c context.Context, taskId string) (Task, error)

}

type TaskUseCase interface {
	Create(c context.Context, task *Task) (Task, error)
	Update(c context.Context, task *Task) (Task, error)
	Delete(c context.Context, taskId string) error
	GetByUserId(c context.Context, userId string) ([]Task, error)
	GetById(c context.Context, taskId string) (Task, error)
}