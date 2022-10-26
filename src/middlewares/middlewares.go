package middlewares

import (
	//"api/src/autenticacao"
	//"api/src/respostas"
	//"log"
	"api/src/autenticacao"
	"fmt"
	"net/http"

	//"net/http"
	"github.com/gin-gonic/gin"
)

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AUTENTICANDO")
	//token := c.Request.Header.Get("Authorization")
	if erro := autenticacao.ValidarToken(c); erro != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro na validação": erro.Error()})
	}
	fmt.Println("AUTENTICAAAADO")
	c.Next()
	}
}

/* 		}
		if len(c.Keys) == 0{
			c.Keys = make(map[string]interface{})
		}
		c.Keys["user"] = token	 */
		

/* 		erro := autenticacao.ValidarToken(c); 
		if erro != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
			"erro na autorização": erro.Error()})
		} */

/* 		if token == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Não há um token",
			})
		}else{
			fmt.Println("AUTENTICAAAADO") }*/
