package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/souzera/dino/schemas"
)

// LISTAR EDUCADORES

func ListarEducadores(contexto *gin.Context) {
	educadors := []schemas.Educador{}

	if err := db.Find(&educadors).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar educadores")
		return
	}

	sendSucess(contexto, "listar-educadores", educadors)
}

type ListarEducadoresResponse struct {
	Message string `json:"message"`
	Data []schemas.EducadorResponse `json:"data"`
}

// CRIAR EDUCADOR

func CriarEducador(contexto *gin.Context) {
	request := CriarEducadorRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarEducadorRequest(request); err != nil{
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	educador := schemas.Educador{
		Nome: request.Nome,
		UsuarioID: request.UsuarioID,
	}

	if err := db.Create(&educador).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar educador")
		return
	}

	sendSucess(contexto, "criar-educador", educador)
}

type CriarEducadorRequest struct {
	Nome string `json:"nome"`
	UsuarioID uuid.UUID `json:"usuario"`
}

type CriarEducadorResponse struct {
	Message string `json:"message"`
	Data schemas.EducadorResponse `json:"data"`
}

func validarCriarEducadorRequest(request CriarEducadorRequest) error {
	if request.Nome == "" {
		return errorParamRequired("nome")
	}

	return nil

}


// BUSCAR EDUCADOR

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