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
	
	err = MigrateAll(db)
	if err != nil {
		logger.Errorf("Erro ao criar as tabelas: %v", err)
		return nil, err
	}

	return db, nil
}

// MIGRANDO TABELAS

func MigrateUsuario(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Usuario{})
}

func MigrateGerencia(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Gerencia{})
}

func MigrateEscola(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Escola{})
}

func MigrateDiretor(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Diretor{})
}

func MigrateEducador(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Educador{})
}

func MigrateAluno(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Aluno{})
}

func MigrateTurma(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Turma{})
}

func MigrateDisciplina(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Disciplina{})
}

func MigrateNota(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Nota{})
}

func MigrateAula(db *gorm.DB) error{
	return db.AutoMigrate(&schemas.Aula{})
}

func MigrateAll(db *gorm.DB) error{
	if err := MigrateUsuario(db); err != nil {
		return err
	}
	if err := MigrateGerencia(db); err != nil {
		return err
	}
	if err := MigrateEscola(db); err != nil {
		return err
	}
	if err := MigrateDiretor(db); err != nil {
		return err
	}
	if err := MigrateEducador(db); err != nil {
		return err
	}
	if err := MigrateAluno(db); err != nil {
		return err
	}
	if err := MigrateTurma(db); err != nil {
		return err
	}
	if err := MigrateDisciplina(db); err != nil {
		return err
	}
	if err := MigrateNota(db); err != nil {
		return err
	}
	if err := MigrateAula(db); err != nil {
		return err
	}
	
	return nil
}