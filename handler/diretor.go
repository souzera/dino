package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarDiretores(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"diretores": "GET lista de diretores",
	})
}

func CriarDiretor(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"diretores": "POST atribuindo cargo a um usuario",
	})
}

func BuscarDiretor(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"diretores": "GET diretor por id",
	})
}

func AtualizarDiretor(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"diretores": "PUT atualizar diretor",
	})
}

func DeletarDiretor(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"diretores": "DELETE deletar diretor",
	})
}

