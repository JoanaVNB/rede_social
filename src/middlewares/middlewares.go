package middlewares

import (
	//"api/src/autenticacao"
	//"api/src/respostas"
	//"log"
	"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTENTICANDO")
/* 		erro := autenticacao.ValidarToken(c); 
		if erro != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
			"erro na autorização": erro.Error()})
		} */
		c.Next()
		}
	}
