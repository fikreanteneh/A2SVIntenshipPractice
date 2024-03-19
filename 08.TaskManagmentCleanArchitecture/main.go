package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }


}