package banco

import(
	//"time"
	_"github.com/go-sql-driver/mysql"
	"api/src/config"
	"database/sql"
	"fmt"

)

func Conectar() (*sql.DB, error){
	db, erro := sql.Open("mysql", config.StringDeConexao)
	if erro != nil{
		return nil, erro
	}

	if erro =db.Ping(); erro != nil{
		db.Close()
		fmt.Println(erro)
		return nil, erro
	}

	return db, nil

}