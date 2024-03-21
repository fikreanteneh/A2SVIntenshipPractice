package controller

import (
	"TaskManger/domain"
	"TaskManger/middleware"
	"TaskManger/models"

	"github.com/gin-gonic/gin"
)


type TaskController struct {
	taskUseCase domain.TaskUseCase
}

func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}


func (t *TaskController) Create(c *gin.Context) {
	// TODO: Error Handling
	// TODO: Authorization
	var task models.TaskCreate
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	username, _ := c.Get("username")
	createdTask, err := t.taskUseCase.Create(c, username.(string),&task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 201, "Task Created Successfully", createdTask)
}

func (t *TaskController) Delete(c *gin.Context) {
		// TODO: Error Handling
	// TODO: Authorization
	taskId := c.Param("id")
	username, _ := c.Get("username")
	deletedTask, err := t.taskUseCase.Delete(c, username.(string), taskId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "Task Deleted Successfully", deletedTask)

}

func (t *TaskController) Update(c *gin.Context) {
		// TODO: Error Handling
	// TODO: Authorization
	taskId := c.Param("id")
	var task models.TaskUpdate
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	username, _ := c.Get("username")
	updatedTask, err := t.taskUseCase.Update(c, username.(string), taskId, &task)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "Task Updated Successfully", updatedTask)

}

func (t *TaskController) GetById(c *gin.Context) {
		// TODO: Error Handling
	// TODO: Authorization
	taskId := c.Param("id")
	username, _ := c.Get("username")
	task, err := t.taskUseCase.GetById(c, username.(string), taskId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "Task Fetched Successfully", task)
}

func (t *TaskController) GetByUserId(c *gin.Context) {
		// TODO: Error Handling
	// TODO: Authorization
	username, _ := c.Get("username")
	tasks, err := t.taskUseCase.GetByUsername(c, username.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	middleware.SuccessResponseHandler(c, 200, "Tasks Fetched Successfully", tasks)
}
