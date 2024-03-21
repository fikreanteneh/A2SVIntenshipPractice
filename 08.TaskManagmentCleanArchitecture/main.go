package main

import (
	"TaskManger/config"
	"TaskManger/router"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    env, _ := config.Load()
    db , _ := config.GetMongoClient(env.DatabaseURL, env.DatabaseName)
    router.Setup(env, 10*time.Second, db, r)
    r.Run("localhost:" + env.Port)
}