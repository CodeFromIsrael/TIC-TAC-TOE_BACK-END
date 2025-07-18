package routes

import (
	"net/http"
	"tic_tac_toe_BACK-END/src/controllers"
)

var routesUser = []Route{
	{
		Url:                   "/usuarios",
		method:                http.MethodPost,
		function:              controllers.CreateUser,
		requireAuthentication: false,
	},
	{
		Url:                   "/login",
		method:                http.MethodPost,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: false,
	},
	{
		Url:                   "/usuarios/{userId}/Atualizar",
		method:                http.MethodGet,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
	{
		Url:                   "/usuarios/{userId}/adicionar-amigos",
		method:                http.MethodPost,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
}
