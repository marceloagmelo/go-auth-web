package handler

import (
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/marceloagmelo/go-auth-web/api"
	"github.com/marceloagmelo/go-auth-web/logger"
	"github.com/marceloagmelo/go-auth-web/models"
	"github.com/marceloagmelo/go-auth-web/utils"
)

var view = template.Must(template.ParseGlob("views/*.html"))
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

const (
	cookieNomeUsuario  = "nomeUsuario"
	cookieSenhaUsuario = "senhaUsuario"
)

//Health testa conexão com o mysql e rabbitmq
func Health(w http.ResponseWriter, r *http.Request) {
	usuarioHealth, err := api.Health()
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro verificar o heal check", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Usuários",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	data := map[string]interface{}{
		"titulo":   "Lista de Usuários",
		"mensagem": usuarioHealth.Mensagem,
	}

	err = view.ExecuteTemplate(w, "Health", data)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro ao chamar a página de health check", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Usuários",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}

//Index primeira página
func Index(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, _ := getDadosCookie(r)
	if !utils.IsEmpty(nomeUsuario) {
		http.Redirect(w, r, "/listar", 302)
	} else {
		data := map[string]interface{}{
			"titulo":   "Login",
			"mensagem": "",
		}

		view.ExecuteTemplate(w, "Login", data)
	}
}

//Listar usuários
func Listar(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, senhaUsuario := getDadosCookie(r)

	if !utils.IsEmpty(nomeUsuario) {
		usuarios, _ := api.ListaUsuarios(nomeUsuario, senhaUsuario)

		data := map[string]interface{}{
			"titulo":      "Lista de Usuários",
			"usuarios":    usuarios,
			"nomeUsuario": nomeUsuario,
		}

		err := view.ExecuteTemplate(w, "Index", data)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro ao chamar a página home", err)
			data := map[string]interface{}{
				"titulo":       "Lista de Usuários",
				"mensagemErro": mensagemErro,
			}

			err := view.ExecuteTemplate(w, "Erro", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	} else {
		avisoUsuarioNaoLogado(w)
	}

}

//Adicionar usuário
func Adicionar(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, senhaUsuario := getDadosCookie(r)

	if !utils.IsEmpty(nomeUsuario) {
		if r.Method == "POST" {

			err := r.ParseForm()
			if err != nil {
				mensagemErro := fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
				data := map[string]interface{}{
					"titulo":       "Lista de Usuários",
					"mensagemErro": mensagemErro,
				}

				err := view.ExecuteTemplate(w, "Erro", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

			usuario := r.FormValue("usuario")
			senha := r.FormValue("senha")
			email := r.FormValue("email")

			if usuario != "" && senha != "" && email != "" {
				var usuarioForm models.Usuario
				usuarioForm.ID = 0
				usuarioForm.Nome = usuario
				usuarioForm.Senha = senha
				usuarioForm.Email = email
				usuarioForm.Status = 1

				usuarioRetorno, err := api.AdicionarUsuario(nomeUsuario, senhaUsuario, usuarioForm)
				if err != nil {
					mensagemErro := fmt.Sprintf("%s: %s", "Erro ao adicionar o usuário", err)
					data := map[string]interface{}{
						"titulo":       "Lista de Usuários",
						"mensagemErro": mensagemErro,
					}

					err = view.ExecuteTemplate(w, "Erro", data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					return
				}

				if usuarioRetorno.ID == 0 {
					mensagemErro := fmt.Sprintf("%s", "Erro ao adicionar o usuário, favor veja o log da api.")
					logger.Erro.Println(mensagemErro)
					data := map[string]interface{}{
						"titulo":       "Lista de Usuários",
						"mensagemErro": mensagemErro,
					}

					err = view.ExecuteTemplate(w, "Erro", data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					return
				}

				mensagem := fmt.Sprintf("Usuário %v adicionado com sucesso!", usuarioRetorno.Nome)
				logger.Info.Println(mensagem)

				http.Redirect(w, r, "/", 301)
			}
		}
	} else {
		avisoUsuarioNaoLogado(w)
	}
}

//Logar usuário
func Logar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
			data := map[string]interface{}{
				"titulo":       "Login Usuário",
				"mensagemErro": mensagemErro,
			}

			err := view.ExecuteTemplate(w, "Erro", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		usuario := r.FormValue("usuario")
		senha := r.FormValue("senha")

		if usuario != "" && senha != "" {
			var usuarioForm models.Usuario
			usuarioForm.Nome = usuario
			senhaSum := sha256.Sum256([]byte(senha))
			senhaHash := fmt.Sprintf("%X", senhaSum)
			usuarioForm.Senha = string(senhaHash)

			usuarioRetorno, err := api.Logar(usuarioForm)
			if err != nil {
				mensagemErro := fmt.Sprintf("%s: %s", "Erro ao logar um usuário", err)
				data := map[string]interface{}{
					"titulo":       "Login",
					"mensagemErro": mensagemErro,
				}

				err = view.ExecuteTemplate(w, "Erro", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				return
			}

			if usuarioRetorno.ID == 0 {
				mensagemErro := fmt.Sprintf("%s", "Usuário ou senha inválidos!")
				data := map[string]interface{}{
					"titulo":       "Login",
					"mensagemErro": mensagemErro,
				}

				err = view.ExecuteTemplate(w, "Erro", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				return
			}

			mensagem := fmt.Sprintf("Usuário %v logado com sucesso!", usuarioRetorno.Nome)
			logger.Info.Println(mensagem)

			setCookie(usuarioRetorno.Nome, usuarioRetorno.Senha, w)

			http.Redirect(w, r, "/", 301)
		}
	}
}

//Logout do usuário
func Logout(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, _ := getDadosCookie(r)

	mensagem := fmt.Sprintf("Usuário %v deslogado com sucesso!", nomeUsuario)
	logger.Info.Println(mensagem)

	clearCookie(cookieNomeUsuario, w)
	clearCookie(cookieSenhaUsuario, w)

	http.Redirect(w, r, "/", 302)
}

//Apagar usuário
func Apagar(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, senhaUsuario := getDadosCookie(r)

	if !utils.IsEmpty(nomeUsuario) {
		vars := mux.Vars(r)
		id := vars["id"]
		err := api.ApagarUsuario(nomeUsuario, senhaUsuario, id)
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
			data := map[string]interface{}{
				"titulo":       "Lista de Usuários",
				"mensagemErro": mensagemErro,
			}

			err := view.ExecuteTemplate(w, "Erro", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		mensagem := fmt.Sprintf("Usuário %v apagado com sucesso!", nomeUsuario)
		logger.Info.Println(mensagem)

		http.Redirect(w, r, "/", http.StatusAccepted)
	} else {
		avisoUsuarioNaoLogado(w)
	}

}

//New página de edição de um novo usuário
func New(w http.ResponseWriter, r *http.Request) {
	nomeUsuario, _ := getDadosCookie(r)

	if !utils.IsEmpty(nomeUsuario) {
		data := map[string]interface{}{
			"titulo":   "Novo Usuário",
			"mensagem": "",
		}

		view.ExecuteTemplate(w, "New", data)
	} else {
		avisoUsuarioNaoLogado(w)
	}
}

//setCookie  setar o cookie
func setCookie(usuario, senha string, response http.ResponseWriter) {
	/*value := map[string]string{
		chave: valor,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  chave,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}*/

	usuarioCookie := http.Cookie{
		Name:    cookieNomeUsuario,
		Value:   usuario,
		Expires: time.Now().Add(time.Duration(1) * time.Hour),
	}
	http.SetCookie(response, &usuarioCookie)

	senhaCookie := http.Cookie{
		Name:    cookieSenhaUsuario,
		Value:   senha,
		Expires: time.Now().Add(time.Duration(1) * time.Hour),
	}
	http.SetCookie(response, &senhaCookie)
}

//clearCookie  limpar o cookie
func clearCookie(chave string, response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   chave,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//getDadosCookie recuperar o usuário no cookie
func getDadosCookie(request *http.Request) (string, string) {
	usuario := ""
	senha := ""

	cookieNome, _ := request.Cookie(cookieNomeUsuario)
	if cookieNome != nil {
		usuario = cookieNome.Value
	}
	cookieSenha, _ := request.Cookie(cookieSenhaUsuario)
	if cookieSenha != nil {
		senha = cookieSenha.Value
	}
	return usuario, senha
}

//avisoUsuarioNaoLogado
func avisoUsuarioNaoLogado(w http.ResponseWriter) {
	mensagemErro := fmt.Sprintf("%s", "Usuário precisa estar logado para acessar essa página!")
	logger.Erro.Println(mensagemErro)
	data := map[string]interface{}{
		"titulo":       "Lista de Usuários",
		"mensagemErro": mensagemErro,
	}

	err := view.ExecuteTemplate(w, "Erro", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return

}
