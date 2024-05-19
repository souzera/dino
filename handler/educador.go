package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListarEducadores(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"educadores": "GET lista de educadores",
	})
}

func CriarEducador(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"educadores": "POST criar educador",
	})
}

func BuscarEducador(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"educadores": "GET educador por id",
	})
}

func AtualizarEducador(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"educadores": "PUT atualizar educador",
	})
}

func DeletarEducador(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"educadores": "DELETE deletar educador",
	})
}