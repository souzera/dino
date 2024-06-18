package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/dino/handler"
)

func initializeRoutes(router *gin.Engine){

	handler.InitializeHandler()
	
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/version", func(contexto *gin.Context){
			contexto.JSON(http.StatusOK, gin.H{
				"version": "1.0.0",
			})
		})
		// USUARIOS
		v1.GET("/usuarios", handler.ListarUsuarios)
		v1.GET("/usuarios/:id", handler.BuscarUsuario)
		v1.POST("/usuarios", handler.CriarUsuario)
		v1.PUT("/usuarios/:id", handler.AtualizarUsuario)
		v1.DELETE("/usuarios/:id", handler.DeletarUsuario)
	}
}