package model

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type CreateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	NewUserID string `json:"newUserId"`
}
