package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConexaoBD() *sql.DB {
	conexao := "user=postgres dbname=nome-do-db host=localhost port=PORT password=password sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
