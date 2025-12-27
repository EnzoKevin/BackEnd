package handlers

import (
	model "BACKEND/internal/models"
	"encoding/json"
	"net/http"
)

func (u Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", u.getAllUsers)
	http.HandleFunc("POST /users", u.addUsers)
}

func ( u Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.useCases.GetAllUsers()
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func ( u Handlers) addUsers(w http.ResponseWriter, r *http.Request) {
	
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}

	id, err := u.useCases.Add(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}

	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.CreateUserResponse{NewUserID: id})
}