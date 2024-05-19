package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarUsuarios(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "GET lista de usuarios",
	})
}

func CriarUsuario(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "GET lista de usuarios",
	})
}

func BuscarUsuario(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "GET usuario por id",
	})
}

func AtualizarUsuario(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "PUT atualizar usuario",
	})
}

func DeletarUsuario(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "DELETE deletar usuario",
	})
}

