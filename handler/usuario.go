package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/souzera/dino/schemas"
)

// LISTAR USUARIOS

func ListarUsuarios(contexto *gin.Context) {

	usuarios := []schemas.Usuario{}

	if err := db.Find(&usuarios).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar usuarios")
		return
	}

	sendSucess(contexto, "listar-usuarios", usuarios)
}

type ListarUsuariosResponse struct {
	Message string `json:"message"`
	Data []schemas.UsuarioResponse `json:"data"`
}

// CRIAR USUARIO

func CriarUsuario(contexto *gin.Context) {
	request := CriarUsuarioRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarUsuarioRequest(request); err != nil{
		logger.Infof("request: %v", request)
		logger.Errorf("Erro ao validar request: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	usuario := schemas.Usuario{
		Login: request.Login, 
		Email: request.Email, 
		Senha: request.Senha,
	}

	if err := db.Create(&usuario).Error; err != nil {
		logger.Errorf("Erro ao criar usuario: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar usuario")
		return
	}

	sendSucess(contexto, "criar-usuario", usuario)
}

type CriarUsuarioRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func validarCriarUsuarioRequest(request CriarUsuarioRequest) error{
	if request.Login == "" {return errorParamRequired("login")}
	if request.Email == "" {return errorParamRequired("email")}
	if request.Senha == "" {return errorParamRequired("senha")}
	
	if query := db.Where("login = ?", request.Login).Or("email = ?", request.Email).First(&schemas.Usuario{}); query != nil {
		// TODO: tratar login e email em linhas separadas
		return errorUniqueViolation("login")
	}

	return nil
}

type CriarUsuarioResponse struct {
	Message string `json:"message"`
	Data schemas.UsuarioResponse `json:"data"`
}


// BUSCAR USUARIO

func BuscarUsuario(contexto *gin.Context) {
	
	usuario := schemas.Usuario{}

	id := contexto.Param("id")

	if err := db.Find(&usuario).Where("id = ?", id).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar usuario")
		return
	}

	sendSucess(contexto, "buscar-usuario", usuario)
}

type BuscarUsuarioRequest struct {
	ID uuid.UUID `json:"id"`
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