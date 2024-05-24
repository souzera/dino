package config

import (
	"os"

	"github.com/souzera/dino/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error){
	logger := GetLogger("DINO")
	
	pathDb := "./db/main.db"
	
	_, err := os.Stat(pathDb)
	if os.IsNotExist(err) {
		logger.Info("Database não encontrada, criando uma nova...")
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Erro ao criar o diretório: %v", err)
			return nil, err
		}

		file, err := os.Create(pathDb)
		if err != nil {
			logger.Errorf("Erro ao criar o arquivo: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(pathDb), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar com o banco de dados: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Usuario{})
	if err != nil {
		logger.Errorf("Erro ao criar as tabelas: %v", err)
		return nil, err
	}

	return db, nil
}