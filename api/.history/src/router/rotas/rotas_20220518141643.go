package rotas

import "net/http"

// Rota representa todas rotas da api.
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
