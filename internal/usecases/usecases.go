package usecases

import (
	model "BACKEND/internal/models"
	"BACKEND/internal/repositories"
	"errors"
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
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}