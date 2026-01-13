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

func (u *UseCases) DeleteUser(id string) error {
	deleted := u.repos.User.DeleteUser(id)
	if !deleted {
		return errors.New("user not found")
	}
	return nil
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
		Password: newUser.Password,
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}