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

func (u *UseCases) GetTrainByID(userID string) (*model.CreateTreino, error) {
	train, err := u.repos.User.GetTreino(userID)
	
	if err != nil {
		return nil, err
	}

	if train == nil {
		return nil, errors.New("train not found")
	}

	return train, nil
}


func (u *UseCases) GetTrainByUserID(id string) (model.CreateTreino, error) {
    var repoReq model.CreateTreino

    user, err := u.repos.User.GetByID(id)
    if err != nil {
        return repoReq, fmt.Errorf("erro ao acessar o banco de dados: %v", err)
    }
    if user == nil {
        return repoReq, fmt.Errorf("usuário com ID %s não encontrado no banco 'goback'", id)
    }

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
- Retorne apenas o JSON, sem explicações ou texto adicional

Formato do JSON:
- Devolva como um JSON separado da seguinte forma:
{
  "dia_1": [
	"agachamento: 3 series de 12 repetições com 60 segundos de descanso",
	"leg press: 3 series de 10 repetições com 60 segundos de descanso"
]}
"dia_2": {
	[
	"supino reto: 3 series de 10 repetições com 60 segundos de descanso",
	"supino inclinado: 3 series de 12 repetições com 60 segundos de descanso"]}
}


IMPORTANTE: O campo "exercicios" deve ser uma lista de STRINGS, não objetos.

o treino deve seguir essa estrutura golang

`

    prompt := fmt.Sprintf(PROMPT_TREINO, user.Height, user.Weight, user.BType, user.Target)

    treinoRaw, err := gemini.GenerateTreino(prompt)
    if err != nil {
        return repoReq, fmt.Errorf("falha ao gerar treino com IA: %v", err)
    }
	fmt.Printf("Treino bruto gerado pela IA: %s\n", treinoRaw)

    treinoRaw = strings.TrimSpace(treinoRaw)
    treinoRaw = strings.TrimPrefix(treinoRaw, "```json")
    treinoRaw = strings.TrimSuffix(treinoRaw, "```")
    treinoRaw = strings.TrimSpace(treinoRaw)

    err = json.Unmarshal([]byte(treinoRaw), &repoReq)
    if err != nil {
        return repoReq, fmt.Errorf("IA retornou JSON inválido: %v | Raw: %s", err, treinoRaw)
    }

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