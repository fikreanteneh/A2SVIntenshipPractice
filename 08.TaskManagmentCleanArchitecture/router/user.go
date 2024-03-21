package router

import (
	"TaskManger/controller"
	"TaskManger/repository"
	"TaskManger/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUseCase(userRepository, timeout)
	userController := controller.NewUserController(userUseCase)
	group.DELETE("/", userController.Delete)
	group.PUT("/updateUsername", userController.UpdateUsername)
	group.GET("/updatePassword", userController.UpdatePassword)
}