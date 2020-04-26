package api

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-auth-web/logger"
	"github.com/marceloagmelo/go-auth-web/utils"
)

//ApiRequest estrutura ApiRequest
type (
	Requisicao interface {
		cmdRequest() (*http.Response, error)
	}
	ApiRequest struct {
		EndPoint     string
		Metodo       string
		NomeUsuario  string
		SenhaUsuario string
		Body         io.Reader
	}
)

//ExecutarRequest request
func ExecutarRequest(req Requisicao) (*http.Response, error) {
	resposta, err := req.cmdRequest()

	return resposta, err
}

//cmdRequest request
func (apiR ApiRequest) cmdRequest() (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest(apiR.Metodo, apiR.EndPoint, apiR.Body)
	if err != nil {
		usuario := fmt.Sprintf("%s: %s", "Erro ao criar um request", err.Error())
		logger.Erro.Println(usuario)
		return nil, err
	}

	if !utils.IsEmpty(apiR.NomeUsuario) {
		request.SetBasicAuth(apiR.NomeUsuario, apiR.SenhaUsuario)
	}

	resposta, err := cliente.Do(request)
	if err != nil {
		usuario := fmt.Sprintf("%s: %s", "Erro ao abrir o request", err.Error())
		logger.Erro.Println(usuario)
		return nil, err
	}
	return resposta, nil
}
