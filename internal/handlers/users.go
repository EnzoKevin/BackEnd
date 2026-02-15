package handlers

import (
	model "BACKEND/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

func (u Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", u.getAllUsers)
	http.HandleFunc("GET /users/{id}", u.getUserByID)
	http.HandleFunc("POST /users", u.addUsers)
	http.HandleFunc("GET /users/{id}/train", u.getTrainByUserID)
	http.HandleFunc("DELETE /users/{id}", u.deleteUser)
	http.HandleFunc("GET /users/{id}/train/UsersTrain", u.getTrainByID)
}

func (u Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.useCases.GetAllUsers()
	
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: "no users found, at handlers"})
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u Handlers) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	user, err := u.useCases.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u Handlers) getTrainByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/users/"):]
	
	segments := strings.Split(path, "/")
	UserId := segments[0]
	
	train, err := u.useCases.GetTrainByID(UserId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(train)

}

func (u Handlers) getTrainByUserID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/users/"):] // Resultado: "8VLy8XAv6TrRG8VIw8Su/train"
    
    // Divide a string pela barra e pega apenas o primeiro elemento (o ID)
    segments := strings.Split(path, "/")
    id := segments[0]
	treino, err := u.useCases.GetTrainByUserID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(treino)
}


func (u Handlers) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	err := u.useCases.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Reason: err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
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