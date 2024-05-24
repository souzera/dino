package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error{
	var err error

	host := "localhost"
	user := "postgres"
	password := "1234"
	dbName := "dino"
	port := "5432"
	sslmode := "disable"
	timeZone := "America/Fortaleza"

	db, err = InitializePostgres(host, user, password, dbName, port, sslmode, timeZone)
	if err != nil {
		return fmt.Errorf("error ao inicializar o banco de dados: %v", err)
	}
	return nil
}

func GetDatabase() *gorm.DB{
	return db
}

func GetLogger(p string) *Logger{
	logger := NewLogger(p)
	return logger
}