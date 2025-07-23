package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"tic_tac_toe_BACK-END/src/autentication"
	"tic_tac_toe_BACK-END/src/banco"
	"tic_tac_toe_BACK-END/src/models"
	"tic_tac_toe_BACK-END/src/repository"
	"tic_tac_toe_BACK-END/src/response"
	"tic_tac_toe_BACK-END/src/segurity"
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
	if erro = user.Prepare("cadastro"); erro != nil {
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

func Login(w http.ResponseWriter, r *http.Request) {
	boryRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		fmt.Println("Erro ao ler o corpo da requisição ")
		return
	}

	var user models.User
	if erro = json.Unmarshal(boryRequest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		fmt.Println("Erro ao converter o json")
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("Erro ao conectar com o banco ")
		return
	}

	defer db.Close()

	repository := repository.NewRepositoryUser(db)

	userSaveInBd, erro := repository.SearchByEmail(user.Email)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("Erro ao buscar o usuário")
		return
	}

	if erro = segurity.Checkpassword(userSaveInBd.Password, user.Password); erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		fmt.Println("Senha incorreta")
		return
	}

	token, erro := autentication.Createtoken(userSaveInBd.Id)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("Erro ao gerar o token")
		return
	}
	userId := strconv.FormatUint(userSaveInBd.Id, 10)

	response.Json(w, http.StatusOK, models.AuthenticationData{ID: userId, Token: token})

}
