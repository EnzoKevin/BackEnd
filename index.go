package main

import (
	"database/sql"
	"log"

	/* 	"BACKEND/controllers"
	 */"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/repositories/user"
	"BACKEND/internal/usecases"

	_ "github.com/microsoft/go-mssqldb"
	/* 	"github.com/gin-gonic/gin"
	 */)

func main() {

connStr := "sqlserver://DESKTOP-DJP8DA5//LOCALDB#E94B6469?database=GO_BACK&trusted_connection=true"

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	userRepo := user.NewUserRepository(db)

	repos := repositories.New(userRepo)

	usecases := usecases.New(repos)

	h := handlers.New(usecases)

	h.Listen(8080)

	/* server := gin.Default()

	server.GET("/treinos", controllers.GetTreinos)

	server.Run("LocalHost:8080")

	println("Hello, World!") */
}