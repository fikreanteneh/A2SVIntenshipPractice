package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


type Task struct {
	id          string    `bson:"_id" json:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      bool      `json:"status"`
}

type CreateTaskType struct {
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    Status      bool      `json:"status"`
}

var client, err = GetMongoClient()
var collection = client.Database("taskmanager").Collection("task")


func GetTask(context *gin.Context) {
	var tasks []Task
	cursor, err := collection.Find(context, bson.M{})
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to get tasks"})
		return
	}
	defer cursor.Close(context)
	for cursor.Next(context) {
		var task Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	context.JSON(200, gin.H{"response": tasks})
}

func GetTaskById(context *gin.Context) {
	var id = context.Param("id")
	var task Task
	err = collection.FindOne(context, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to get task"})
		return
	}
	context.JSON(200, gin.H{"response": task})
}

func CreateTask(context *gin.Context) {
	var task CreateTaskType
    if err := context.BindJSON(&task); err != nil {
        context.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
	var cursor, err = collection.InsertOne(context, bson.M{"title": task.Title, "description": task.Description, "status": task.Status, "dueDate": task.DueDate})
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to create task"})
		return
	}
	context.JSON(201, gin.H{"result": cursor, "message": "Successfully Created Task"})
}


func UpdateTask(context *gin.Context) {
	var id = context.Param("id")
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
	updateResult, err := collection.UpdateOne(context, bson.M{"_id": id}, bson.M{"$set": bson.M{"title": task.Title, "description": task.Description, "status": task.Status, "dueDate": task.DueDate}})
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to update task"})
		return
	}
	context.JSON(200, gin.H{"response": updateResult, "message": "Successfully Updated Task"})
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
	context.JSON(200, gin.H{
		"reponse":deleteResult,"message": "Successfully deleted task"})
}
