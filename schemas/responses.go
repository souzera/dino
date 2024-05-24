package schemas

import (
	"time"
)

type UsuarioResponse struct {
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Login	string `json:"login"`
	Email	string `json:"email"`
	Senha 	string `json:"-"`
}

type GerenciaResponse struct {}

type EscolaResponse struct {}

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
	ID		string	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Usuario UsuarioResponse `json:"usuario"`
	Nome string `json:"nome"`
}