package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarGre(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"gre": "GET lista de gre",
	})
}

func CriarGre(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"gre": "POST criar gre",
	})
}

func BuscarGre(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"gre": "GET gre por id",
	})
}

func AtualizarGre(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"gre": "PUT atualizar gre",
	})
}

func DeletarGre(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"gre": "DELETE deletar gre",
	})
}
