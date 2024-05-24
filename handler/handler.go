package handler

import (
	"github.com/souzera/dino/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("DINO")
	db = config.GetDatabase()
}