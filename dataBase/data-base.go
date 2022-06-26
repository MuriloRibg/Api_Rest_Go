package dataBase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConectaMySQL() *sql.DB {
	stringDeConexao := fmt.Sprintf("%s:%s@/%s", "root", "@Murilo1202", "CURSOS")

	db, err := sql.Open("mysql", stringDeConexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
