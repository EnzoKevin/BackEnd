package main

import (
	"BACKEND/controllers"
	"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	repos := repositories.New()

	usecases := usecases.New(repos)

	h := handlers.New(usecases)

h.Listen(8080)

	server := gin.Default()

	server.GET("/treinos", controllers.GetTreinos)

	server.Run("LocalHost:8080")

	println("Hello, World!")
}