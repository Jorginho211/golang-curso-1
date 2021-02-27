package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoes,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoesPorUsuario,
		RequerAuthenticacao: true,
	},
}
