package controllers

import(
	"github.com/gin-gonic/gin"
	"api/src/models"
	"api/src/banco"
	"api/src/repositorio"
	"net/http"
	"strconv"
)

func CriarUsuario(c *gin.Context){
	var usuario models.Usuario

//vincular usuario no content
	if err := c.ShouldBindJSON(&usuario); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if deuRuim := usuario.Preparar("cadastro"); deuRuim != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": deuRuim.Error()})
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
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao criar usuário": erro.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"Inserido": usuarioID})
}

func BuscarUsuarios(c *gin.Context){

	nome := c.Param("nome") 
	
	db, erro := banco.Conectar()
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao conectar com banco de dados": erro.Error()})
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nome)
	if erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao buscar usuario": erro.Error()})
	}
	c.JSON(200, usuarios)
}

func BuscarUsuario(c *gin.Context){
	
	id, erro := strconv.ParseUint(c.Param("id"), 10, 64)
	if erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao converter para uint": erro.Error()})
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
	usuario, erro := repositorio.BuscarPorID(id)
	if erro != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"erro ao buscar ID": erro.Error()})
	}
	c.JSON(200, usuario)
}

func AtualizarUsuario(c *gin.Context){
	var usuario models.Usuario

	id, erro := strconv.ParseUint(c.Param("id"), 10, 64)
	if erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao converter para uint": erro.Error()})
		return
	}
	//lê o corpo da requisição
	if err := c.ShouldBindJSON(&usuario); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error()})
	}

	db, erro := banco.Conectar()
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao conectar com banco de dados": erro.Error()})
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
	if erro := repositorio.Atualizar(id, usuario); erro !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error()})
	}
	c.JSON(200, gin.H{
			"message": "Usuário atualizado",
		}) 
}
	
func DeletarUsuario(c *gin.Context){
	id, erro := strconv.ParseUint(c.Param("id"), 10, 64)
	if erro != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao converter para uint": erro.Error()})
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
	if erro := repositorio.Deletar(id); erro !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error()})
	}
	c.JSON(200, gin.H{
		"message": "Usuário Deletado",
	})
}