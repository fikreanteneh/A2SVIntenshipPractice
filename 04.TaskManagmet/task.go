package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


type Task struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	DueDate     time.Time   `json:"dueDate"`
	Status      bool        `json:"status"`
}

type CreateTaskType struct {
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    Status      bool      `json:"status"`
}

var db = make(map[int]Task)
var dbid = 1;

func GetTask(context *gin.Context) {
	context.JSON(200, db)
}

func GetTaskById(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	var _, ok = db[id]
		fmt.Println(id, ok)

	if err != nil || !ok {
		context.JSON(400, "Invalid Id")
		return
	}
	context.JSON(200, db[id])
}

func CreateTask(context *gin.Context) {
	var task CreateTaskType
    if err := context.Bind(&task); err != nil {
        context.JSON(400, "Invalid Request Type")
        return
    }
	var created = Task{
		Id: dbid,
		Title: task.Title,
		Description: task.Description,
		Status: task.Status,
		DueDate: task.DueDate,
	}
	db[created.Id] = created
	dbid++
	context.JSON(201, "Successfully Added task")
}


func UpdateTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	var _, ok = db[id]	

	if err != nil || !ok {
		context.JSON(400, "Invalid Id")
		return
	}
	var task CreateTaskType
	context.BindJSON(&task)
	db[id] = Task{
		Id: id,
		Title: task.Title,
		Description: task.Description,
		Status: task.Status,
		DueDate: task.DueDate,
	}
	context.JSON(200, "Successfully Updated Task")
}

func DeleteTask(context *gin.Context) {
	var id, err = strconv.Atoi(context.Param("id"))
	var _, ok = db[id]
	fmt.Println(id, ok)
	if err != nil || !ok {
		context.JSON(400, "Invalid Id")
		return
	}
	delete(db, id)
	context.JSON(200, db)
}
