package model

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Weight   float64
	Height   float64
	BType    string
	Target   string
}

type CreateUserRequest struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	BType    string  `json:"btype"`
	Target   string  `json:"target"`
}

type CreateUserResponse struct {
	NewUserID string `json:"newUserId"`
}
type CreateTreino struct {
	ID      string   `json:"id"`
	Segunda []string `json:"dia_1"`
	Terca   []string `json:"dia_2"`
	Quarta  []string `json:"dia_3"`
	Quinta  []string `json:"dia_4"`
	Sexta   []string `json:"dia_5"`
	Sabado  []string `json:"dia_6"`
	Domingo []string `json:"dia_7"`
}