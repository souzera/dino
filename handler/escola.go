package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarEscolas(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"escolas": "GET lista de escolas",
	})
}

func CriarEscola(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"escolas": "POST criar escola",
	})
}

func BuscarEscola(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"escolas": "GET escola por id",
	})
}

func AtualizarEscola(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"escolas": "PUT atualizar escola",
	})
}

func DeletarEscola(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"escolas": "DELETE deletar escola",
	})
}