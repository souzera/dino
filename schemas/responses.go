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