package routes

import "net/http"

var onllineRoutes = []Route{
	{
		Url:                   "/onlline",
		method:                http.MethodGet,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
	{
		Url:                   "/onlline/{userId}/solicitar-partida-amigo",
		method:                http.MethodPost,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
	{
		Url:                   "/onlline/{userId}/solicitar-partida",
		method:                http.MethodPost,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
	{
		Url:                   "/onlline/{userId}/aceitar-partida",
		method:                http.MethodPost,
		function:              func(w http.ResponseWriter, r *http.Request) {},
		requireAuthentication: true,
	},
}
