package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Login	string `gorm:"unique" json:"login"`
	Email	string `json:"email"`
	Senha	string `json:"-"`
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	OwnerType string `json:"owner_type,omitempty"`
}

type Gerencia struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Nome string `json:"nome"`
	Escolas []Escola `gorm:"foreignKey:GerenciaID" json:"escolas"`
}

type Escola struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Nome	string `json:"nome"`
	Endereco string `json:"endereco"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Gerencia Gerencia `gorm:"default:null" json:"gerencia"`
	GerenciaID uuid.UUID `json:"gerencia_id,omitempty"`
	Diretor Diretor `gorm:"default:null" json:"diretor"`
	DiretorID uuid.UUID `json:"diretor_id,omitempty"`
}

type Turma struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Nome string `json:"nome"`
	Escola Escola `gorm:"default:null" json:"escola"`
	EscolaID uuid.UUID `json:"escola_id,omitempty"`
	Alunos []Aluno `gorm:"many2many:aluno_turma;" json:"alunos"`
}

type Diretor struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Usuario Usuario `gorm:"polymorphic:Owner;" json:"usuario"`
	Nome string `json:"nome"`
}

type Educador struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Usuario Usuario `gorm:"polymorphic:Owner;" json:"usuario"`
	Nome string `json:"nome"`
}

type Aluno struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Usuario Usuario `gorm:"polymorphic:Owner;" json:"usuario"`
	Nome string `json:"nome"`
}

type Disciplina struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Nome string `json:"nome"`
	Turma Turma `gorm:"default:null" json:"turma"`
	TurmaID uuid.UUID `json:"turma_id,omitempty"`
}

type Nota struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Aluno Aluno `gorm:"default:null" json:"aluno"`
	AlunoID uuid.UUID `json:"aluno_id,omitempty"`
	Disciplina Disciplina `gorm:"default:null" json:"disciplina"`
	DisciplinaID uuid.UUID `json:"disciplina_id,omitempty"`
	Ativdade float64 `json:"atividade"`
	Bimestre int `json:"bimestre"`
}

type Aula struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	Disciplina Disciplina `gorm:"default:null" json:"disciplina"`
	DisciplinaID uuid.UUID `json:"disciplina_id,omitempty"`
	Data time.Time `json:"data"`
	Conteudo string `json:"conteudo"`
	Frequencia []Aluno `gorm:"many2many:frequencia" json:"frequencia"`
}