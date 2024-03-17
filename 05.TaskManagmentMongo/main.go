package main

import (
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.GET("/tasks", GetTask)
	router.GET("/tasks/:id", GetTaskById)
	router.POST("/tasks", CreateTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)
	router.Run("localhost:5555")
}