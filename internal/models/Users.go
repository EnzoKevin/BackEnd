package model

type User struct {
	ID       string  `json:"id" firestore:"-"` // O "-" diz: "Não procure isso nos campos do banco"
	Name     string  `json:"name" firestore:"name"`
	Email    string  `json:"email" firestore:"email"`
	Password string  `json:"password" firestore:"password"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	BType    string  `json:"btype"`
	Target   string  `json:"target"`
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