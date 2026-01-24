package usecases

import (
	"BACKEND/internal/gemini"
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"

	"errors"
	"fmt"
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

func (u *UseCases) GetTrainByID(id string) (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	user, err := u.repos.User.GetByID(id)
	
	
	prompt := fmt.Sprintf(os.Getenv("PROMPT_TREINO"), user.Height, user.Weight, user.BType, user.Target)
	
	treino, err := gemini.GenerateTreino(prompt)
	if err != nil {
		panic(err)
	}
	var repoReq model.CreateTreino
	repoReq.ID = user.ID

	err = json.Unmarshal([]byte(treino), &repoReq)
	if err != nil {
		return "", "Falha na conversão linha 67", fmt.Errorf("erro ao converter treino: %v", err)
	}



	AddTreino, err := u.repos.User.AddTreino(repoReq)

	return AddTreino, treino, err
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