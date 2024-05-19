package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/version", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"version": "1.0.0",
			})
		})
		// USUARIOS
		v1.GET("/usuarios", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"usuarios": "GET lista de usuarios",
			})
		})
		v1.GET("/usuarios/:id", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"usuarios": "GET usuario por id",
			})
		})
		v1.POST("/usuarios", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"usuarios": "POST criar usuario",
			})
		})
		v1.PUT("/usuarios/:id", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"usuarios": "PUT atualizar usuario",
			})
		})
		v1.DELETE("/usuarios/:id", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"usuarios": "DELETE deletar usuario",
			})
		})
	}
}