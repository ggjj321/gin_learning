package main

import (
	"EP2/database"
	"EP2/middleares"
	"EP2/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogging()

	router := gin.Default()

	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middleares.Logger())

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8000")
}
