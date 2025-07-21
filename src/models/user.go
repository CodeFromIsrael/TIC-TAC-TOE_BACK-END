package models

import (
	"errors"
	"strings"
	"tic_tac_toe_BACK-END/src/segurity"

	"github.com/badoux/checkmail"
)

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Name     string `json:"nome,omitempty"`
	Phone    string `json:"telefone,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"senha,omitempty"`
	Onlline  bool   `json:"onlline,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.trimFields(step); err != nil {
		return err
	}
	if err := user.validate(); err != nil {
		return err
	}
	return nil
}

func (user *User) trimFields(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	if step == "cadastro" {
		passHash, erro := segurity.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passHash)
	}
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O campo nome é obrigatório")
	}
	if user.Phone == "" {
		return errors.New("O campo telefone é obrigatório")
	}
	if user.Email == "" {
		return errors.New("O campo e-mail é obrigatório")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail é inválido")
	}
	if user.Password == "" {
		return errors.New("O campo senha é obrigatório")
	}
	return nil
}
