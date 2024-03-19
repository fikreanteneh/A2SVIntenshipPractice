package controller

import (
	"TaskManger/domain"

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
	panic("unimplemented")
}

func (t *TaskController) Delete(c *gin.Context) {

}

func (t *TaskController) Update(c *gin.Context) {

}

func (t *TaskController) GetById(c *gin.Context) {

}

func (t *TaskController) GetByUserId(c *gin.Context) {

}
