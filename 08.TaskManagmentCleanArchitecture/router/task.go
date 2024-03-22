package router

import (
	"TaskManger/config"
	"TaskManger/controller"
	"TaskManger/repository"
	"TaskManger/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db, "task")
	userRepository := repository.NewUserRepository(db, "user")
	taskUseCase := usecase.NewTaskUseCase(taskRepository, userRepository, environment, timeout)
	taskController := controller.NewTaskController(taskUseCase)
	group.POST("/", taskController.Create)
	group.DELETE("/:id", taskController.Delete)
	group.PUT("/:id", taskController.Update)
	group.GET("/:id", taskController.GetById)
	group.GET("/", taskController.GetByUserId)
}