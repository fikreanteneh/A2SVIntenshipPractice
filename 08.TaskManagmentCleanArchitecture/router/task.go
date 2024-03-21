package router

import (
	"TaskManger/controller"
	"TaskManger/repository"
	"TaskManger/usecase"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db, "task")
	userRepository := repository.NewUserRepository(db, "user")
	taskUseCase := usecase.NewTaskUseCase(taskRepository, userRepository, timeout)
	taskController := controller.NewTaskController(taskUseCase)
	group.POST("/", taskController.Create)
	group.DELETE("/:taskId", taskController.Delete)
	group.PUT("/:taskId", taskController.Update)
	group.GET("/:taskId", taskController.GetById)
	group.GET("/", taskController.GetByUserId)
}