package domain

import (
	"TaskManger/models"
	"context"
	"time"
)


type Task struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      bool   `json:"status" bson:"status"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
	UserId      string `json:"userId" bson:"userId"`
}


type TaskRepository interface {
	Create(c context.Context, task *Task) (*Task, error)
	Update(c context.Context, task *Task) (*Task, error)
	Delete(c context.Context, task *Task) (*Task, error)
	GetByUserId(c context.Context, userId string) (*[]*Task, error)
	GetById(c context.Context, taskId string) (*Task, error)
}

type TaskUseCase interface {
	Create(c context.Context, username string, payload *models.TaskCreate) (*Task, error)
	Update(c context.Context, username string,  taskId string, payload *models.TaskUpdate) (*Task, error)
	Delete(c context.Context, username string,taskId string) (*Task, error)
	GetByUsername(c context.Context, username string) (*[]*Task, error)
	GetById(c context.Context, username string, tasId string) (*Task, error)
}