package main

import (
	"fmt"
	"net/http"
	"tic_tac_toe_BACK-END/src/config"
	"tic_tac_toe_BACK-END/src/router"
)

func main() {
	config.Toload()
	r := router.Generete()
	http.ListenAndServe(fmt.Sprintf(":%d", config.Door), r)
}
