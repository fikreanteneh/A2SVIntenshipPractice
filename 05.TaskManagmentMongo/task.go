package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Task struct {
	id          primitive.ObjectID  `bson:"_id"`         `json:"_id,omitempty"`
	Title       string              `bson:"title"`       `json:"title"`
	Description string              `bson:"description"` `json:"description"`
	DueDate     time.Time           `bson:"dueDate"`     `json:"dueDate"`
	Status      bool                `bson:"status"`      `json:"status"`
}


type CreateTaskType struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	DueDate     time.Time   `json:"dueDate"`
	Status      bool        `json:"status"`

}

var db = make(map[int]Task)
var dbid = 1;
client, err := GetMongoClient()
collection := client.Database("taskmanager").Collection("task")

func GetTask(context *gin.Context) {
	var tasks []Task
	if err != nil {
		return context.JSON(400, "Failed to get tasks")
	}
	cursor, err := collection.Find(context, bson.M{})
	if err != nil {
		return context.JSON(400, "Failed to get tasks")
	}
	defer cursor.Close(context)
	for cursor.Next(context) {
		var task Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return context.JSON(200, tasks)
}

func GetTaskById(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	var _, ok = db[id]
	if err != nil || !ok {
		context.JSON(400, gin.H{"error": "Invalid Id"})
		return
	}
	var task Task
	err = collection.FindOne(context, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to get task"})
		return
	}
	context.JSON(200, task)
}

func CreateTask(context *gin.Context) {
	var task CreateTaskType
    if err := context.BindJSON(&task); err != nil {
        context.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

	
	var cursor, err = collection.InsertOne(context, bson.M{"title": task.title, "description": task.description, "status": task.status, "dueDate": task.dueDate})
    if err != nil {
        return context.JSON(400, gin.H{"error": "Failed to create task"})
    }
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to create task"})
		return
	}
	context.JSON(201, gin.H{"result": cursor, "message": "Successfully Created Task"})
}


func UpdateTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil{
		context.JSON(400, gin.H{"error": "Invalid Id"})
		return
	}
	var task CreateTaskType
	context.BindJSON(&task)
	if err != nil {
		context.JSON(400, "Failed to update task")
		return
	}
	updateResult, err := collection.UpdateOne(context, bson.M{"_id": id}, bson.M{"$set": bson.M{"title": task.title, "description": task.description, "status": task.status, "dueDate": task.dueDate}})
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}
	context.JSON(200, gin.H{"result": updateResult, "message": "Successfully Updated Task"})
}

func DeleteTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}
	deleteResult, err := collection.DeleteOne(context, bson.M{"_id": id})
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to delete task"})
		return
	}
	context.JSON(200, gin.H{"result", "Successfully deleted task"})
}
