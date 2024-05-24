package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/dino/schemas"
)

func ListarUsuarios(contexto *gin.Context) {
	contexto.JSON(http.StatusOK, gin.H{
		"usuarios": "GET lista de usuarios",
	})
}

func CriarUsuario(contexto *gin.Context) {
	request := struct{
		Login string `json:"login"`
		Email string `json:"email"`
		Senha string `json:"senha"`
	}{}

	contexto.BindJSON(&request)


	usuario := schemas.Usuario{Login: request.Login, Email: request.Email, Senha: request.Senha}

	logger.Infof("Criando usuario: %v", request.Email)

	if err := db.Create(&usuario).Error; err != nil {
		logger.Errorf("Erro ao criar usuario: %v", err)
		return
	}

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

