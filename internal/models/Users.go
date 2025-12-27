package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserRequest struct {
	Username string    `json:"username"`
}

type CreateUserResponse struct {
	NewUserID uuid.UUID `json:"newUserId"`
}