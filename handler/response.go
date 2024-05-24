package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(contexto *gin.Context, code int, mensagem string) {
	contexto.Header("Content-Type", "application/json")
	contexto.JSON(code, gin.H{
		"message":   mensagem,
		"errorCode": code,
	})
}

func sendSucess(contexto *gin.Context, op string, data interface{}) {
	contexto.Header("Content-Type", "application/json")
	contexto.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s realizado com sucesso", op),
		"data":    data,
	})
}