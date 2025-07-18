package banco

import (
	"database/sql"
	"tic_tac_toe_BACK-END/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConectionDB)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil

}
