package repositories

import (
	model "BACKEND/internal/models"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Add(model.User) (string, error)        
	AddTreino(model.CreateTreino) (string, error)
	GetByID(id string) (*model.User, error)
	DeleteUser(id string) error           
	EmailExists(email string) (bool, error)
	GetTreino(UserID string) (*model.CreateTreino, error)
}

type Repositories struct {
	User UserRepository
}

func New(repo UserRepository) *Repositories {
	return &Repositories{
		User: repo,
	}
}
