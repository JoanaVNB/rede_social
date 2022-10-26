package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorio"
	"api/src/seguranca"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"strconv"
)

func Login(c *gin.Context){
	//usuario que esta vindo na requisição
	var usuario models.Usuario
	//ler o corpo da requisição
	if err := c.ShouldBindJSON(&usuario); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao conectar com banco de dados": erro.Error()})
		return
	}
	defer db.Close()


	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
//usuario que esta salvo no banco
usuarioSalvo, erro := repositorio.BuscarPorEmail(usuario.Email)
if erro != nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"erro ao buscar por e-mail": erro.Error()	})
	}
	if erro = seguranca.VerificarSenha(usuarioSalvo.Senha, usuario.Senha); erro != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro ao verificar senha": erro.Error()})
	}
	fmt.Println("Autorizado")
	token, erro:= autenticacao.CriarToken(usuarioSalvo.ID)
	if erro != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro ao criar Token": erro.Error()})
	}
	fmt.Printf("Token: %s", token)
}
