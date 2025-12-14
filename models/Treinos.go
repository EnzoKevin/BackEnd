package model

type Treinos struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      uint   `json:"userId"`
}