package autenticacao

import (
	"api/src/config"
	//"errors"
	"fmt"
	//"strconv"
	"strings"
	"time"
	//"net/http"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()//expirar em 6 horas
	permissoes["usuarioId"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)//um tipo de método de assinatura de token e permissões
	return token.SignedString([]byte(config.SecretKey)) //criação da assinatura com a chave SecretKey
}

// ValidarToken verifica se o token passado na requisição é valido
func ValidarToken(c *gin.Context) error {
	
	tokenString := extrairToken(c)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)//vai estar num formato que de pra ler
	if erro != nil {
		return erro
	}
	
	fmt.Println(token)
	return nil
/* 	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido") */
}

func extrairToken(c *gin.Context) string{
	token := c.GetHeader("Authorization") 
	//token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}