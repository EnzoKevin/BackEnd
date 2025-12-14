package main

import (
	"BACKEND/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/treinos", controllers.GetTreinos)

	server.Run("LocalHost:8080")

	println("Hello, World!")
}