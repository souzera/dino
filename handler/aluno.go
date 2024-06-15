package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/souzera/dino/schemas"
)

func ListarAlunos(contexto *gin.Context) {

	alunos := []schemas.Aluno{}

	if err := db.Find(&alunos).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar alunos")
		return
	}

	sendSucess(contexto, "listar-alunos", alunos)
}

type ListarAlunosResponse struct {
	Message string                  `json:"message"`
	Data    []schemas.AlunoResponse `json:"data"`
}

func CriarAluno(contexto *gin.Context) {
	request := CriarAlunoRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarAlunoRequest(request); err != nil {
		logger.Infof("request: %v", request)
		logger.Errorf("Erro ao validar request: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	aluno := schemas.Aluno{
		Nome:      request.Nome,
		Ano:       request.Ano,
		UsuarioID: request.UsuarioID,
	}

	if err := db.Create(&aluno).Error; err != nil {
		logger.Errorf("Erro ao criar aluno: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar aluno")
		return
	}

	sendSucess(contexto, "criar-aluno", aluno)
}

func validarCriarAlunoRequest(request CriarAlunoRequest) error {
	if request.Nome == "" {
		return errorParamRequired("Nome")
	}

	if request.UsuarioID == uuid.Nil{
		return errorParamRequired("Usuario")
	}

	return nil
}

type CriarAlunoRequest struct {
	Nome      string `json:"nome"`
	UsuarioID uuid.UUID `json:"usuario_id"`
	Ano       int    `json:"ano"`
}
