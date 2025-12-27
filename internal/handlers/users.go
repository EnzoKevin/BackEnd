package handlers

import (
	model "BACKEND/internal/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (u Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", u.getAllUsers)
	http.HandleFunc("POST /users", u.addUsers)
}

func ( u Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]model.User{})
}

func ( u Handlers) addUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.CreateUserResponse{NewUserID: uuid.New()})
}