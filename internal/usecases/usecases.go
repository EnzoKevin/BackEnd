package usecases

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories"
	"errors"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{repos: repos}
}

func (u *UseCases) GetAllUsers() []model.User {
	users, _ := u.repos.User.GetAll()
	return users
}

func (u UseCases) Add(newUser model.CreateUserRequest) (uuid.UUID, error) {
	exists, err := u.repos.User.EmailExists(newUser.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if exists {
		return uuid.Nil, errors.New("user already exists")
	}

	repoReq := model.User{
		ID: uuid.New(),
		Name: newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}