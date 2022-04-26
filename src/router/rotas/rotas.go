package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota represents all the API routes
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// Configurar insert all the routes into the router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasFunds

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao)
	}
	return r
}
