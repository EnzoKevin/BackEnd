package repositories

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories/user"
	"database/sql"
)

type Repositories struct {
	User interface {
		GetAll() ([]model.User, error)
		Add(user model.User) (int, error)
		EmailExists(email string) (bool, error)
		GetUserById(id int) (*model.User, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
	}
}
