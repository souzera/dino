package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/dino/schemas"
)

func ListarEscolas(contexto *gin.Context) {
	
	escolas := []schemas.Escola{}

	if err := db.Find(&escolas).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar escolas")
		return
	}

	sendSucess(contexto, "listar-escolas", escolas)
}

type ListarEscolasResponse struct {
	Message string `json:"message"`
	Data []schemas.EscolaResponse `json:"data"`
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