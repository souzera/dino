package main

import (
	"github.com/souzera/dino/config"
	"github.com/souzera/dino/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("DINO")

	err := config.Init()
	if err != nil {
		logger.Errorf("Erro ao inicializar a aplicação: %v", err)
		return
	}
	router.Initialize()
}