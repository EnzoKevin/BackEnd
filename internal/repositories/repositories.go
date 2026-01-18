package repositories

import (
	model "BACKEND/internal/models"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Add(model.User) (string, error)        // 🔥 string
	GetByID(id string) (*model.User, error)
	DeleteUser(id string) error             // 🔥 error
	EmailExists(email string) (bool, error)
}

type Repositories struct {
	User UserRepository
}

func New(repo UserRepository) *Repositories {
	return &Repositories{
		User: repo,
	}
}
