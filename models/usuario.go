package models

//Usuario estrutura de usuário
type Usuario struct {
	ID     int    `db:"id" json:"id"`
	Login  string `db:"login" json:"login"`
	Senha  string `db:"senha" json:"senha"`
	Email  string `db:"email" json:"email"`
	Status int    `db:"status" json:"status"`
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
