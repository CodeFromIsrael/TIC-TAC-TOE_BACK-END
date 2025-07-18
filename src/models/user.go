package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Name     string `json:"nome,omitempty"`
	Phone    string `json:"telefone,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"senhas,omitempty"`
	Onlline  bool   `json:"onlline,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.trimFields(); err != nil {
		return err
	}
	if err := user.validate(); err != nil {
		return err
	}
	return nil
}

func (user *User) trimFields() error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
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
