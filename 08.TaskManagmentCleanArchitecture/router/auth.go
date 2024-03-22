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

func NewAuthRouter(environment *config.Environment, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db, "user")
	userUseCase := usecase.NewUserUseCase(userRepository, environment, timeout)
	userController := controller.NewUserController(userUseCase)
	group.POST("/register", userController.Register)
	group.POST("/login", userController.Login)
}