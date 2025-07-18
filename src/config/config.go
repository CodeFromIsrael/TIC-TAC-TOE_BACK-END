package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConectionDB = ""

	Door = 0
)

func Toload() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Door, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Door = 9000
	}
	StringConectionDB = fmt.Sprintf("%s:%s@/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_BANK"))

}
