package api

import (
	"bytes"
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
	var apiRequest ApiRequest
	apiRequest.EndPoint = variaveis.ApiURL + "/" + api + "/health"
	apiRequest.Metodo = "GET"

	resposta, err := ExecutarRequest(apiRequest)
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
func ListaUsuarios(nomeUsuario, senhaUsuario string) (usuarios models.Usuarios, erro error) {
	var apiRequest ApiRequest
	apiRequest.EndPoint = variaveis.ApiURL + "/" + api + "/usuarios"
	apiRequest.NomeUsuario = nomeUsuario
	apiRequest.SenhaUsuario = senhaUsuario
	apiRequest.Metodo = "GET"

	resposta, err := ExecutarRequest(apiRequest)
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
func AdicionarUsuario(nomeUsuario, senhaUsuario string, novoUsuario models.Usuario) (usuarioRetorno models.Usuario, erro error) {
	var apiRequest ApiRequest
	apiRequest.EndPoint = variaveis.ApiURL + "/" + api + "/usuario/adicionar"
	apiRequest.NomeUsuario = nomeUsuario
	apiRequest.SenhaUsuario = senhaUsuario
	apiRequest.Metodo = "POST"

	conteudoEnviar, err := json.Marshal(&novoUsuario)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro ao gerar o objeto com o JSON lido", err.Error())
		logger.Erro.Println(mensagemErro)
		return usuarioRetorno, err
	}
	apiRequest.Body = bytes.NewBuffer(conteudoEnviar)

	resposta, err := ExecutarRequest(apiRequest)
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
	var apiRequest ApiRequest
	apiRequest.EndPoint = variaveis.ApiURL + "/" + api + "/usuario/login"
	apiRequest.Metodo = "POST"

	conteudoEnviar, err := json.Marshal(&usuario)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro ao gerar o objeto com o JSON lido", err.Error())
		logger.Erro.Println(mensagemErro)
		return usuarioRetorno, err
	}
	apiRequest.Body = bytes.NewBuffer(conteudoEnviar)

	resposta, err := ExecutarRequest(apiRequest)
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
func ApagarUsuario(nomeUsuario, senhaUsuario string, id string) error {
	var apiRequest ApiRequest
	apiRequest.EndPoint = variaveis.ApiURL + "/" + api + "/usuario/apagar/" + id
	apiRequest.NomeUsuario = nomeUsuario
	apiRequest.SenhaUsuario = senhaUsuario
	apiRequest.Metodo = "DELETE"

	resposta, err := ExecutarRequest(apiRequest)
	if err != nil {
		return err
	}
	defer resposta.Body.Close()

	return nil
}
