package router

import (
	"api/src/controllers"
	"api/src/middlewares"
	//"net/http"
	"github.com/gin-gonic/gin"
)


func HandleRequests() {
	api := gin.Default()
	secured := api.Group("/secured") //localhost:8080/secured/login
	{
		secured.Use(middlewares.Autenticar())//middleware ->aplica uma função para todas as rotas
		secured.POST("/login", controllers.Login)
	}
	api.POST("/usuarios", controllers.CriarUsuario)
	api.GET("/usuarios/id/:id", controllers.BuscarUsuario)
	api.PUT("/usuarios/id/:id", controllers.AtualizarUsuario)
	api.DELETE("/usuarios/id/:id", controllers.DeletarUsuario)
		
	api.Run()
	
	// no curso, a requisição de autorização é no mesmo grupo
}
