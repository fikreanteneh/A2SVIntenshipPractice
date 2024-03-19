package router

import (
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"TaskManger/middleware"
	"TaskManger/config"
)

func Setup(env config.Environment, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("auth")

	userRouter := gin.Group("user")
	userRouter.Use(middleware.AuthMiddleware(env.JwtSecret))

	taskRouter := gin.Group("task")
	taskRouter.Use(middleware.AuthMiddleware(env.JwtSecret))


	NewAuthRouter(env, timeout, db, publicRouter)
	NewUserRouter(env, timeout, db, userRouter)
	NewTaskRouter(env, timeout, db, taskRouter)
}
