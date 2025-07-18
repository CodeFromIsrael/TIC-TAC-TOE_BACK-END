package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"tic_tac_toe_BACK-END/src/banco"
	"tic_tac_toe_BACK-END/src/models"
	"tic_tac_toe_BACK-END/src/repository"
	"tic_tac_toe_BACK-END/src/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	boryRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var user models.User
	if erro = json.Unmarshal(boryRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = user.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Connect()
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	user.Id, erro = repository.Create(user)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.Json(w, http.StatusCreated, user)
}
