package models

import "time"

//Usuario estrutura de usuário
type Usuario struct {
	ID              int       `db:"id" json:"id"`
	Nome            string    `db:"nome" json:"nome"`
	Senha           string    `db:"senha" json:"senha"`
	Email           string    `db:"email" json:"email"`
	DataCriacao     time.Time `db:"dtcriacao" json:"dtcriacao"`
	DataAtualizacao time.Time `db:"dtatualizacao" json:"dtatualizacao"`
	Status          int       `db:"status" json:"status"`
}

//Usuarios lista de usuarios
type Usuarios []Usuario

//UsuarioHealth retorno do health
type UsuarioHealth struct {
	Mensagem string `json:"mensagem"`
}

//MensagemErro retorno do erro
type MensagemErro struct {
	Erro string `json:"erro"`
}
