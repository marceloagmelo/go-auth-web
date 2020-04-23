package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marceloagmelo/go-auth-web/logger"
	"github.com/marceloagmelo/go-auth-web/models"
	"github.com/marceloagmelo/go-auth-web/variaveis"
)

var api = "go-auth/api/v1"

//Health testar conexão com a API
func Health() (usuarioHealth models.UsuarioHealth, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/health"

	resposta, err := GetRequest(endpoint)
	if err != nil {
		return usuarioHealth, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioHealth, err
		}
		usuarioHealth = models.UsuarioHealth{}
		err = json.Unmarshal(corpo, &usuarioHealth)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioHealth, err
		}
	}
	return usuarioHealth, nil
}

//ListaUsuarios listar mensagens
func ListaUsuarios() (usuarios models.Usuarios, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/usuarios"

	resposta, err := GetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
		usuarios = models.Usuarios{}
		err = json.Unmarshal(corpo, &usuarios)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
	}
	return usuarios, nil
}

//AdicionarUsuario adicionar usuário
func AdicionarUsuario(novoUsuario models.Usuario) (usuarioRetorno models.Usuario, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/usuario/adicionar"

	resposta, err := PostRequest(endpoint, novoUsuario)
	if err != nil {
		return usuarioRetorno, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == http.StatusCreated {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioRetorno, err
		}
		usuarioRetorno = models.Usuario{}
		err = json.Unmarshal(corpo, &usuarioRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioRetorno, err
		}
	}
	return usuarioRetorno, nil
}

//Logar logar com o usuário
func Logar(usuario models.Usuario) (usuarioRetorno models.Usuario, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/usuario/login"

	resposta, err := PostRequest(endpoint, usuario)
	if err != nil {
		return usuarioRetorno, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == http.StatusOK {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioRetorno, err
		}
		usuarioRetorno = models.Usuario{}
		err = json.Unmarshal(corpo, &usuarioRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return usuarioRetorno, err
		}
	}
	return usuarioRetorno, nil
}

//ApagarUsuario apagar usuário
func ApagarUsuario(id string) error {
	endpoint := variaveis.ApiURL + "/" + api + "/usuario/apagar/" + id

	err := DeleteRequest(endpoint)
	if err != nil {
		return err
	}
	return nil
}
