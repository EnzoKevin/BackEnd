package user

import (
	model "BACKEND/internal/models"
)

type Users struct {
	users []model.User
}

func New() *Users {
	return &Users{users: make([]model.User, 0)}
}

func (u Users) GetAll() []model.User {
	return u.users
}

func (u *Users) Add(NewUser model.User) {
	u.users = append(u.users, NewUser)
}