package repository

import (
	"database/sql"
	"errors"
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

func (User *User) SearchByEmail(email string) (models.User, error) {
	lines, erro := User.db.Query("select id,senha from usuarios where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(&user.Id, &user.Password); erro != nil {
			return models.User{}, erro
		}
		return user, nil
	}
	return user, errors.New("usuário não encontrado ")
}
