package repositorio

import (
	"database/sql"
	"api/src/models"
	"fmt"

)

type Usuarios struct {
	db *sql.DB
}

//cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios{
	return &Usuarios{db}
}
//método vai estar dentro da struct Usuarios 
//Insere um usuário no banco de dados
func (u Usuarios) Criar(usuario models.Usuario) (uint64, error){
	statement, erro := u.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",)
	if erro != nil{
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil{
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}
 
func (u Usuarios) Buscar(nome string) ([]models.Usuario, error){
	nome = fmt.Sprintf("%%%s%%", nome) // %nome%

	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ?",
		nome,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (u Usuarios) BuscarPorID (ID uint64) (models.Usuario, error){
	linhas, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	//usuarios com valor zero:
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro//struct em branco e erro
		}
	}
	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (u Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	statement, erro := u.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
} 

// Deletar exclui as informações de um usuário no banco de dados
func (u Usuarios) Deletar(ID uint64) error {
	statement, erro := u.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

func (u Usuarios) BuscarPorEmail (email string) (models.Usuario, error){
	linha, erro := u.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {//preencher id e senha
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}