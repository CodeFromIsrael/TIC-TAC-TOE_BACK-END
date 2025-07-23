package middlewares

import (
	"log"
	"net/http"
	"tic_tac_toe_BACK-END/src/autentication"
	"tic_tac_toe_BACK-END/src/response"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Autentication(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autentication.ValidadeToken(r); erro != nil {
			response.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}
