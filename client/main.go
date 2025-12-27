package main

import (
	model "BACKEND/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	req := model.CreateUserRequest{
		Name: "testuser",
		Email: "test@gmail.com",
	}

	b, err :=json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err :=http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	fmt.Println("STATUS:", resp.Status)

	var responseAPI model.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseAPI); err != nil {
		panic(err)
	}

	fmt.Println("New User ID:", responseAPI.NewUserID)
}