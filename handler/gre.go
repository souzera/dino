package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/dino/schemas"
)

func ListarGre(contexto *gin.Context) {
	
	gres := []schemas.Gerencia{}

	if err := db.Find(&gres).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar gres")
		return
	}

	sendSucess(contexto, "listar-gres", gres)
}

type ListarGreResponse struct {
	Message string `json:"message"`
	Data []schemas.GerenciaResponse `json:"data"`
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
