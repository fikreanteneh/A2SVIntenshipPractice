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

func NewUserRouter(environment *config.Environment,timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUseCase(userRepository, environment, timeout)
	userController := controller.NewUserController(userUseCase)
	group.DELETE("/", userController.Delete)
	group.PUT("/updateUsername", userController.UpdateUsername)
	group.GET("/updatePassword", userController.UpdatePassword)
}