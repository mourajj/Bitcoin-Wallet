package rotas

import (
	"maxxer/src/controllers"
	"net/http"
)

var rotasFunds = []Rota{
	{
		URI:    "/deposit",
		Metodo: http.MethodPost,
		Funcao: controllers.Deposit,
	},
	{
		URI:    "/withdraw",
		Metodo: http.MethodPut,
		Funcao: controllers.Withdraw,
	},
	{
		URI:    "/balance/{user}",
		Metodo: http.MethodGet,
		Funcao: controllers.Balance,
	},
	{
		URI:    "/history/{user}",
		Metodo: http.MethodGet,
		Funcao: controllers.History,
	},
	{
		URI:    "/richestperson",
		Metodo: http.MethodGet,
		Funcao: controllers.RichestPerson,
	},
}
