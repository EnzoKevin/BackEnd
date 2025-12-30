package model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID 
	Name string    
	Email    string   
	Password string  
}

type CreateUserRequest struct {
	Name string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserResponse struct {
	NewUserID uuid.UUID `json:"newUserId"`
}

