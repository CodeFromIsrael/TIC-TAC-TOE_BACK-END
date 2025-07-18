package router

import (
	"tic_tac_toe_BACK-END/src/router/routes"

	"github.com/gorilla/mux"
)

func Generete() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
