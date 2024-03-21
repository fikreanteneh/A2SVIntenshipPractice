package repository

import (
	"TaskManger/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	database   *mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository {
	return &TaskRepository{
		database:   db,
		collection: collection,
	}
}


// Create implements domain.TaskRepository.
func (t *TaskRepository) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
	task.Id = primitive.NewObjectID().Hex()
	result, err := t.database.Collection(t.collection).InsertOne(c, task)
	return &domain.Task{
		Id:          result.InsertedID.(string),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
		UserId:      task.UserId,
	}, err
}

// Delete implements domain.TaskRepository.
func (t *TaskRepository) Delete(c context.Context, task *domain.Task) (*domain.Task, error) {
	filter := bson.M{"_id": task.Id}
	_, err := t.database.Collection(t.collection).DeleteOne(c, filter)
	return task, err
}

// GetById implements domain.TaskRepository.
func (t *TaskRepository) GetById(c context.Context, taskId string) (*domain.Task, error) {
	filter := bson.M{"_id": taskId}
	result := t.database.Collection(t.collection).FindOne(c, filter)
	var task domain.Task
	err := result.Decode(&task);
	return &task, err
}

// GetByUserId implements domain.TaskRepository.
func (t *TaskRepository) GetByUserId(c context.Context, userId string) (*[]*domain.Task, error) {
    filter := bson.M{"userId": userId}
    cursor, err := t.database.Collection(t.collection).Find(c, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(c)
    var tasks []*domain.Task
    if err = cursor.All(c, &tasks); err != nil {
        return nil, err
    }
    return &tasks, nil
}

// Update implements domain.TaskRepository.
func (t *TaskRepository) Update(c context.Context, task *domain.Task) (*domain.Task, error) {
	    filter := bson.M{"_id": task.Id}
    update := bson.M{"$set": task}

    _, err := t.database.Collection(t.collection).UpdateOne(c, filter, update)
    if err != nil {
        return nil, err
    }

    return task, nil
	
}


