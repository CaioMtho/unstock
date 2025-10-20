package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "../../db/unstock_db")
	if err != nil {
		log.Fatal("Erro ao conectar ao banco: " + err.Error())
	}
}