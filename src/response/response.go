package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, statuscode int, erro error) {
	Json(w, statuscode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
