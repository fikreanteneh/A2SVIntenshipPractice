package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/register")
	router.POST("/login")
	router.DELETE("/deleteAccount")
	router.PUT("/updateAccount")
	router.Run("localhost:5555")

}