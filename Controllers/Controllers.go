package controllers

import (
	model "BACKEND/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var treino = []model.Treinos{
	{ID: 1, Name: "Treino A", Description: "Treino para membros superiores", UserId: 1},
	{ID: 2, Name: "Treino B", Description: "Treino para membros inferiores", UserId: 1},
}

func GetTreinos(c *gin.Context) { 
	c.IndentedJSON(http.StatusOK, treino)
}