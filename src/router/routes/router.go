package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url                   string
	method                string
	function              func(http.ResponseWriter, *http.Request)
	requireAuthentication bool
}

func Config(r *mux.Router) *mux.Router {
	route := routesUser

	for _, rota := range route {
		r.HandleFunc(rota.Url, rota.function).Methods(rota.method)
	}
	return r
}
