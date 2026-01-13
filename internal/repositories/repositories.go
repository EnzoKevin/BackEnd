package repositories

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories/user"
)

type Repositories struct {
	User interface {
		GetAll() ([]model.User, error)
		Add(model.User) (int, error)
		GetByID(id string) (*model.User, error)
		DeleteUser(id string) bool
		EmailExists(email string) (bool, error)
	}
}

func New(repo *user.UserRepo) *Repositories {
	return &Repositories{
		User: repo,
	}
}
