package main

import (
	"log"

	"BACKEND/internal/DB"
	"BACKEND/internal/gemini"
	"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/repositories/user"
	"BACKEND/internal/usecases"
	"fmt"
)

func main() {

altura := 175
	peso := 78
	tipoCorporal := "mesomorfo"
	objetivo := "hipertrofia"

	prompt := fmt.Sprintf(`
Você é um personal trainer profissional.

Crie um treino semanal com base nos seguintes dados:

Altura: %d cm
Peso: %d kg
Tipo corporal: %s
Objetivo: %s

Regras:
- Divida o treino por dias da semana
- Informe exercícios, séries, repetições e descanso
- Linguagem clara e objetiva
- Não use termos médicos complexos
`, altura, peso, tipoCorporal, objetivo)

	treino, err := gemini.GenerateTreino(prompt)
	if err != nil {
		panic(err)
	}

	fmt.Println("🏋️ Treino gerado pela Gemini:\n")
	fmt.Println(treino)

firebaseDB, err := DB.ConnectDB()
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

	