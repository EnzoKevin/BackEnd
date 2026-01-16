package main

import (
	"log"

	"BACKEND/internal/db"
	"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/repositories/user"
	"BACKEND/internal/usecases"
)

func main() {

firebaseDB, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewUserRepository(firebaseDB.Firestore)

	repos := repositories.New(userRepo)

	usecases := usecases.New(repos)

	h := handlers.New(usecases)

	h.Listen(8080)

	/* server := gin.Default()

	server.GET("/treinos", controllers.GetTreinos)

	server.Run("LocalHost:8080")

	println("Hello, World!") */
}

	