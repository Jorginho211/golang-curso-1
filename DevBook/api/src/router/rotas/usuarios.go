package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuarios,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuario,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/usuario/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequerAuthenticacao: false,
	},
}
