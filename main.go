package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"time"
)

func conectaMySQL() *sql.DB {
	stringDeConexao := fmt.Sprintf("%s:%s@/%s", "root", "@Murilo1202", "CURSOS")

	db, err := sql.Open("mysql", stringDeConexao)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*html")) //pegando as páginas.

func main() {
	db := conectaMySQL()
	defer db.Close()
	http.HandleFunc("/", index) //resebe o caminho e função q será executada.

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaMySQL()
	selectProduto, err := db.Query("SELECT * FROM PRODUTOS;")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &quantidade, &nome, &descricao, &preco)
		if err != nil {
			panic(err.Error())
		}
	}

	temp.ExecuteTemplate(w, "index", produtos)
}
