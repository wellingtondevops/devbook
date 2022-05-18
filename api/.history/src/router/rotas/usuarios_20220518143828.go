package rotas

import (

	"net/http"
)

// slice de rotas de usuarios

var rotasUsuarios = []Rota{

	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:            
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:            
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:            
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:            
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             
		RequerAutenticacao: false,
	},
}
