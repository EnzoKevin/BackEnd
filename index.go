package main

import (
	/* 	"BACKEND/controllers"
	 */"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/usecases"
	"database/sql"
	"log"
	/* 	"github.com/gin-gonic/gin"
	 */)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=backend sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	repos := repositories.New(db)

	usecases := usecases.New(repos)

	h := handlers.New(usecases)

	h.Listen(8080)

	/* server := gin.Default()

	server.GET("/treinos", controllers.GetTreinos)

	server.Run("LocalHost:8080")

	println("Hello, World!") */
}