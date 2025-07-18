package repository

import (
	"database/sql"
	"tic_tac_toe_BACK-END/src/models"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {
	return &User{db}
}

func (user *User) Create(u models.User) (uint64, error) {
	smt, erro := user.db.Prepare("insert into  usuarios (nome,telefone,email,senha) values(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}

	defer smt.Close()

	insert, erro := smt.Exec(u.Name, u.Phone, u.Email, u.Password)
	if erro != nil {
		return 0, erro
	}
	lastId, erro := insert.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastId), erro
}
