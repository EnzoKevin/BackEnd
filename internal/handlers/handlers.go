package handlers

import (
	"BACKEND/internal/usecases"
	"fmt"
	"net/http"
)

type Handlers struct {
	useCases usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	
	
	return &Handlers{useCases: *useCases}
}

func (h Handlers) Listen (port int) error {
	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}

