package config

import (
	"fmt"

	"github.com/souzera/dino/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres(host string, user string, password string, dbName string, port string, sslmode string, timeZone string) (*gorm.DB, error){

	logger := GetLogger("DINO")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslmode, timeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar com o banco de dados: %v", err)
		return nil, err
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	
	err = db.AutoMigrate(&schemas.Usuario{}, &schemas.Gerencia{}, &schemas.Escola{}, &schemas.Diretor{}, &schemas.Educador{}, &schemas.Aluno{}, &schemas.Turma{}, &schemas.Disciplina{}, &schemas.Nota{}, &schemas.Aula{})

	if err != nil {
		logger.Errorf("Erro ao criar as tabelas: %v", err)
		return nil, err
	}

	return db, nil
}