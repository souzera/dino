package schemas

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Login	string `json:"login"`
	Email	string `json:"email"`
	Senha	string `json:"-"`
}