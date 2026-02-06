package usecases

import (
	"BACKEND/internal/gemini"
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories"
	"encoding/json"

	"errors"
	"fmt"
	"strings"
)



type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{repos: repos}
}

func (u *UseCases) GetAllUsers() []model.User {
	users, err := u.repos.User.GetAll()
	if err != nil {
		return nil
	}
	return users
}

func (u *UseCases) GetUserByID(id string) (*model.User, error) {
	user, err := u.repos.User.GetByID(id)
	
	if err != nil {
		return nil, err
	}
	
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

/* func (u *UseCases) GetTrainByID(id string) (string, error) {
	
var PROMPT_TREINO = `Você é um personal trainer profissional.

Crie um treino semanal com base nos seguintes dados:

Altura: %f cm
Peso: %f kg
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
	"supino inclinado": "3 series de 12 repetições com 60 segundos de descanso"]}
}



o treino deve seguir essa estrutura golang

`
	

	user, err := u.repos.User.GetByID(id)
	
	
	prompt := fmt.Sprintf(PROMPT_TREINO, user.Height, user.Weight, user.BType, user.Target)
	
	treino, err := gemini.GenerateTreino(prompt)
	if err != nil {
		panic(err)
	}
	var repoReq model.CreateTreino
	repoReq.ID = user.ID

	err = json.Unmarshal([]byte(treino), &repoReq)
	if err != nil {
		return  "Falha na conversão linha 67", fmt.Errorf("erro ao converter treino: %v", err)
	}



	AddTreino, err := u.repos.User.AddTreino(repoReq)

	return AddTreino, err
} */

func (u *UseCases) GetTrainByID(id string) (model.CreateTreino, error) {
    var repoReq model.CreateTreino

    // 1. Busca o usuário e evita o NIL POINTER DEREFERENCE
    user, err := u.repos.User.GetByID(id)
    if err != nil {
        return repoReq, fmt.Errorf("erro ao acessar o banco de dados: %v", err)
    }
    if user == nil {
        // Se o banco 'goback' não encontrar o ID, retornamos erro em vez de dar panic
        return repoReq, fmt.Errorf("usuário com ID %s não encontrado no banco 'goback'", id)
    }

    // 2. Prompt com chaves alinhadas à sua struct CreateTreino
    var PROMPT_TREINO = `Você é um personal trainer profissional.

Crie um treino semanal com base nos seguintes dados:

Altura: %f cm
Peso: %f kg
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
	"supino inclinado": "3 series de 12 repetições com 60 segundos de descanso"]}
}



o treino deve seguir essa estrutura golang

`

    prompt := fmt.Sprintf(PROMPT_TREINO, user.Height, user.Weight, user.BType, user.Target)

    // 3. Chamada ao Gemini
    treinoRaw, err := gemini.GenerateTreino(prompt)
    if err != nil {
        return repoReq, fmt.Errorf("falha ao gerar treino com IA: %v", err)
    }

    // 4. LIMPEZA DO JSON (O Gemini costuma envolver o JSON em ```json ... ```)
    treinoRaw = strings.TrimSpace(treinoRaw)
    treinoRaw = strings.TrimPrefix(treinoRaw, "```json")
    treinoRaw = strings.TrimSuffix(treinoRaw, "```")
    treinoRaw = strings.TrimSpace(treinoRaw)

    // 5. Conversão para a Struct
    err = json.Unmarshal([]byte(treinoRaw), &repoReq)
    if err != nil {
        return repoReq, fmt.Errorf("IA retornou JSON inválido: %v | Raw: %s", err, treinoRaw)
    }

    // 6. Vincula o ID do usuário ao treino e salva
    repoReq.ID = user.ID
    _, err = u.repos.User.AddTreino(repoReq)
    if err != nil {
        return repoReq, fmt.Errorf("erro ao persistir treino no Firestore: %v", err)
    }

    return repoReq, nil
}

func (u *UseCases) DeleteUser(id string) error {
	err := u.repos.User.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (u UseCases) Add(newUser model.CreateUserRequest) (string, error) {
	exists, err := u.repos.User.EmailExists(newUser.Email)

	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("user already exists")
	}

	repoReq := model.User{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
		Password: newUser.Password,
		Weight: newUser.Weight,
		Height: newUser.Height,
		BType: newUser.BType,
		Target: newUser.Target,
		
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}