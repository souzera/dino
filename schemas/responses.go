package schemas

import (
	"time"

	"github.com/google/uuid"
)

type UsuarioResponse struct {
	ID		uuid.UUID	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Login	string `json:"login"`
	Email	string `json:"email"`
	Senha 	string `json:"-"`
}

type GerenciaResponse struct {
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Nome string `json:"nome"`
	Escolas []EscolaResponse `json:"escolas"`
}

type EscolaResponse struct {
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Nome	string `json:"nome"`
	Endereco string `json:"endereco"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Gerencia GerenciaResponse `json:"gerencia"`
	Diretor DiretorResponse `json:"diretor"`
}

type TurmaResponse struct {}

type DiretorResponse struct {
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Nome string `json:"nome"`
	Usuario UsuarioResponse `json:"usuario"`
}

type EducadorResponse struct {
	ID		uuid.UUID	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Usuario UsuarioResponse `json:"usuario"`
	Nome string `json:"nome"`
}

type AlunoResponse struct {
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Nome string `json:"nome"`
	Usuario UsuarioResponse `json:"usuario"`
	Ano int `json:"ano"`
}