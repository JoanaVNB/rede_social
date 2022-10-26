package autenticacao

import (
	"api/src/config"
	"errors"
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
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { 
		fmt.Println("TOKEN VÁLIDO")
		return nil
	}
	return errors.New("token inválido")
}

func extrairToken(c *gin.Context) string{
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 { //pq retorna "BearerToken chave"
		return strings.Split(token, " ")[1]
	}
	return ""
}

//verifica se o método de assinatura é o que esperamos
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil //se for o método esperado, retorna a chave de verificação
}