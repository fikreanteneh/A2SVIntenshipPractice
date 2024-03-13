package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


type Task struct {
	id          int         `json:id`
	title       string      `json:"title"`
	description string      `json:"description"`
	dueDate     time.Time   `json:"dueDate"`
	status      bool        `json:"status"`
}

var db = make(map[int]Task)

func GetTask(context *gin.Context) {
	context.JSON(200, db)
}

func GetTaskById(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	context.JSON(200, db[id])
}

func CreateTask(context *gin.Context) {
	var task Task
	context.BindJSON(&task)
	task.id = len(db)
	db[task.id] = task
	context.JSON(200, db)
}

func UpdateTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var task Task
	context.BindJSON(&task)
	db[id] = task
	context.JSON(200, db)
}

func DeleteTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	delete(db, id)
	context.JSON(200, db)
}
