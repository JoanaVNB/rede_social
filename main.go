package main

import (
	"api/src/config"
	"api/src/router"
	//"crypto/rand"
	//"log"
	//"encoding/base64"
	"fmt"
	//"api/src/banco"
	//"github.com/gin-gonic/gin"
)

/* func init(){
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro !=nil{
		log.Fatal(erro)
	}
 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Printf("SecretKey: %s", stringBase64)
} */
func init(){
	fmt.Printf("SecretKey: %s/n", config.SecretKey)
} 

func main (){
	config.Carregar()
	//banco.Conectar()
	router.HandleRequests()

}