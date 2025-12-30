package main

import (
	model "BACKEND/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	req := model.CreateUserRequest{
		Name: "testuser",
		Email: "test@gmail.com",
		Password: "teste",
	}

	b, err :=json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err :=http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	
	defer resp.Body.Close()

	fmt.Println("STATUS:", resp.Status)

	// 🔴 NUNCA decodifique se não for sucesso
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		panic(fmt.Sprintf(
			"request failed | status=%d | body=%s",
			resp.StatusCode,
			string(body),
		))
	}
	var responseAPI model.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseAPI); err != nil {
		panic(err)
	}

	fmt.Println("New User ID:", responseAPI.NewUserID)
}