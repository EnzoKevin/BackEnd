package main

import (
	"log"

	"BACKEND/internal/DB"
	"BACKEND/internal/handlers"
	"BACKEND/internal/repositories"
	"BACKEND/internal/repositories/user"
	"BACKEND/internal/usecases"
)

func main() {

/* altura := 175
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
- Devolva como um JSON separado da seguinte forma:
{
  "dia_da_semana_1_PERNAS": {
	"exercicios": [
	"agachamento": "3 series de 12 repetições com 60 segundos de descanso",
	"leg press": "3 series de 10 repetições com 60 segundos de descanso"
]}
"dia_da_semana_2_PEITO": {
	"exercicios": [
	"supino reto": "3 series de 10 repetições com 60 segundos de descanso",
	"supino inclinado": "3 series de 12 repetições com 60 segundos de descanso"]}}
]}
}`, altura, peso, tipoCorporal, objetivo)

	treino, err := gemini.GenerateTreino(prompt)
	if err != nil {
		panic(err)
	}

	fmt.Println("🏋️ Treino gerado pela Gemini:\n")
	fmt.Println(treino) */

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

	