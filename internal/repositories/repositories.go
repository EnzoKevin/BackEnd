package repositories

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories/user"
)

type Repositories struct {
	User interface {
		GetAll() []model.User
		Add(NewUser model.User)
	}
}

func New() *Repositories {
	return &Repositories{
		User: user.New(),
	}
}