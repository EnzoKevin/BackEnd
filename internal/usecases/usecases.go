package usecases

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{repos: repos}
}

func (u *UseCases) GetAllUsers() []model.User {
	users := u.repos.User.GetAll()
	return users
}

func (u UseCases) Add(newUser model.User) uuid.UUID {
	repoReq := model.User{
		ID: uuid.New(),
		Username: newUser.Username,
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID
}