package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct{
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error{
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	return nil
}

//Each call to New returns a distinct error value even if the text is identical.
func (usuario *Usuario) validar(etapa string) error{
/* 	if usuario.Nome == "" || usuario.Nick == "" || usuario.Email == "" || usuario.Senha == ""{
		return errors.New("Favor prencher todos os campos.")
	} */

	if usuario.Nome == ""{
		return errors.New("Favor prencher o campo: nome.")
	}

	if usuario.Nick == ""{
		return errors.New("Favor prencher o campo: nick.")
	}

	if usuario.Email == ""{
		return errors.New("Favor prencher o campo: email.")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil{
		return errors.New("O e-mail inserido não é válido.")
	}

	if etapa == "cadastro" && usuario.Senha == ""{
		return errors.New("Favor prencher o campo: senha.")
	}	
	return nil
}

func (usuario *Usuario) formatar(etapa string) error{
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro"{
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil{
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}